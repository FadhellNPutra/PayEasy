package controller

import (
	"net/http"
	"payeasy/config"
	"payeasy/entity/dto"
	"payeasy/shared/common"
	"payeasy/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUc usecase.AuthUseCase
	rg     *gin.RouterGroup
}

func (a *AuthController) loginHandler(ctx *gin.Context) {
	var payload dto.AuthRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	rsv, err := a.authUc.Login(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, rsv, "Ok")
}

func (a *AuthController) Route() {
	a.rg.POST(config.AuthLogin, a.loginHandler)
}

func NewAuthController(authUc usecase.AuthUseCase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{authUc: authUc, rg: rg}
}
