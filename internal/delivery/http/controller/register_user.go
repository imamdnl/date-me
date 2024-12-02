package controller

import (
	"date-me/internal/delivery/http/mapper"
	"date-me/internal/delivery/http/request"
	"date-me/internal/delivery/http/response"
	"date-me/pkg/config"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func (h *Controller) RegisterUser(c *gin.Context) {
	var payload request.UserRegisterRequest

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

	user, err := mapper.NewMapper().RegisterRequestToEntity(payload)
	if err != nil {
		config.Logger().Error("error mapping request to entity", zap.Error(err))
		response.DecisionResponse(c, err, nil)
		return
	}

	resp := h.userUsecase.RegisterUser(c, user)
	if !resp {
		response.DecisionResponse(c, errors.New("user register failed"), nil)
		return
	}

	response.DecisionResponse(c, nil, resp)
}
