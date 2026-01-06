package controller

import (
	"net/http"

	"github.com/Albaihaqi354/koda-b5-backend/internal/dto"
	"github.com/Albaihaqi354/koda-b5-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (a *AuthController) PostAuthRegister(ctx *gin.Context) {
	var register dto.User
	if err := ctx.ShouldBindJSON(&register); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if errMessage := service.ValidateAuthRegister(register.Email, register.Password); errMessage != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errMessage,
		})
		return
	}

	service.SaveUser(register)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register berhasil",
		"data":    register,
	})
}

func (a *AuthController) PostAuthLogin(ctx *gin.Context) {
	var login dto.User
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if errMessage := service.ValidateAuthRegister(login.Email, login.Password); errMessage != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errMessage,
		})
		return
	}

	user, errMessage := service.LoginUser(login.Email, login.Password)
	if errMessage != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errMessage,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"data":    user,
	})
}
