package router

import (
	"github.com/Albaihaqi354/koda-b5-backend/internal/controller"
	"github.com/gin-gonic/gin"
)

func Init(app *gin.Engine) {
	authControler := controller.NewAuthController()

	app.POST("/auth/register", authControler.PostAuthRegister)
	app.POST("/auth/login", authControler.PostAuthLogin)
}
