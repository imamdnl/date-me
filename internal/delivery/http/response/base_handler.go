package response

import (
	"date-me/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DecisionResponse(ctx *gin.Context, err error, results interface{}) {
	if err != nil {
		utils.RespondWithError(ctx.Writer, http.StatusBadRequest, err)
		return
	} else {
		ctx.AbortWithStatusJSON(http.StatusOK, utils.NewSuccessJsonResponse(results, "SUCCESS"))
		return
	}
}
