package middleware

import (
	"date-me/pkg/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ExceptionMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				errM := fmt.Sprintf("%v", err)
				errMessage := errors.New("Something went wrong on server. " + errM)
				c.AbortWithStatusJSON(http.StatusInternalServerError, utils.NewFailedJsonResponse(errMessage, errMessage.Error()))
			}
		}()
		c.Next()
	})
}
