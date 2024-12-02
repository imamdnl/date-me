package delivery

import (
	"date-me/domain/usecase"
	"date-me/internal/delivery/http/controller"

	"github.com/gin-gonic/gin"
)

type (
	Constructor struct {
		Gin            *gin.Engine
		UserController *controller.Controller
	}
)

func NewRoute(router *gin.Engine, uc *usecase.Wrapper) {
	userController := controller.NewController(*uc)

	mobileGroup := router.Group("/mobile")

	mobileGroup.POST("/v1/register", userController.RegisterUser)
	mobileGroup.POST("/v1/login", userController.LoginUser)
}
