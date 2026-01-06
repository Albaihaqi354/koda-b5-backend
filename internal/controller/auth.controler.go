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

	if !service.IsValidEmail(register.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email tidak valid",
		})
		return
	}

	if len(register.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Password minimal 6 karakter",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register berhasil",
		"data":    register,
	})
}
