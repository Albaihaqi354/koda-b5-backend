package main

import (
	"github.com/Albaihaqi354/koda-b5-backend/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router.Init(app)
	app.Run("localhost:5000")
}
