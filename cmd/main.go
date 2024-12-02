package main

import (
	"context"
	"date-me/internal/delivery"
	"date-me/internal/delivery/middleware"
	"date-me/internal/repository"
	"date-me/internal/usecase"
	"date-me/pkg/common"
	"date-me/pkg/config"
	"errors"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type ConnectionMainApp struct {
	Ctx      context.Context
	BaseRepo *common.BaseRepository
	Logger   *zap.Logger
}

type serverOptions struct {
	Address     string
	Handler     http.Handler
	QuitSignals []os.Signal
}

func main() {
	// initialize service configuration
	config.Environment()

	connMain := new(ConnectionMainApp)
	connMain.Ctx = context.Background()
	connMain.Logger = config.Logger()

	router := connMain.SetupRouter()
	router.Use(middleware.ExceptionMiddleware())

	serverOpt := &serverOptions{
		Address:     ":" + os.Getenv("PORT"),
		Handler:     router,
		QuitSignals: []os.Signal{syscall.SIGINT, syscall.SIGTERM},
	}

	_, err := InitHttpServer(serverOpt)
	if err != nil {
		config.Logger().Fatal("Error starting the server:", zap.Error(err))
	}
}

func InitHttpServer(opt *serverOptions) (*http.Server, error) {
	srv := &http.Server{
		Addr:    opt.Address,
		Handler: opt.Handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			config.Logger().Fatal("listen:", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, opt.QuitSignals...)
	<-quit
	config.Logger().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		config.Logger().Fatal("Server Shutdown:", zap.Error(err))
	}

	select {
	case <-ctx.Done():
		config.Logger().Info("Timeout of 1 seconds.")
	}
	config.Logger().Info("Server exiting")

	return srv, nil
}

func (conn *ConnectionMainApp) SetupRouter() *gin.Engine {
	conn.setupDatabaseConnection()
	router := conn.setupConfigRouter()
	conn.setupServices(router)

	return router
}

func (conn *ConnectionMainApp) setupDatabaseConnection() {
	// initialize repository
	base := common.NewBaseRepository(
		config.PostgreSql(conn.Ctx, conn.Logger),
		nil,
	)

	if err := base.PostgresSql.Ping(conn.Ctx); err != nil {
		panic(err)
	}

	conn.BaseRepo = base
	config.Logger().Info("Database Connection")
}

func (conn *ConnectionMainApp) setupConfigRouter() *gin.Engine {
	router := gin.Default()
	router.RemoveExtraSlash = true
	router.MaxMultipartMemory = 10 << 20 // 10 MiB

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	return router
}

func (conn *ConnectionMainApp) setupServices(router *gin.Engine) {
	var (
		wrapperRepo = repository.Init(conn.BaseRepo, conn.Logger)
		uc          = usecase.Init(wrapperRepo, conn.Logger)
	)

	delivery.NewRoute(router, uc)
}
