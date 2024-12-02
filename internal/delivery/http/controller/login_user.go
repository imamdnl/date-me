package controller

import (
	"date-me/internal/delivery/http/mapper"
	"date-me/internal/delivery/http/request"
	"date-me/internal/delivery/http/response"
	"date-me/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func (h *Controller) LoginUser(c *gin.Context) {
	var payload request.UserLoginRequest

	if err := c.Bind(&payload); err != nil {
		config.Logger().Error("error parsing request body:", zap.Error(err))
		response.DecisionResponse(c, err, nil)
		return
	}

	if invalid := validator.New().Struct(payload); invalid != nil {
		config.Logger().Error("error validate request", zap.Error(invalid))
		response.DecisionResponse(c, invalid, nil)
		return
	}

	userDTO := mapper.NewMapper().LoginRequestToDTO(payload)
	resp, err := h.userUsecase.LoginUser(c, userDTO)
	if err != nil {
		response.DecisionResponse(c, err, nil)
		return
	}

	response.DecisionResponse(c, nil, mapper.NewMapper().UserEntityToLoginResponse(resp))
}
