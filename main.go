package main

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	app := gin.Default()

	app.POST("/auth/register", func(ctx *gin.Context) {
		var register User
		if err := ctx.ShouldBindJSON(&register); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request body",
			})
			return
		}

		if !isValidEmail(register.Email) {
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
	})

	app.Run("localhost:5000")
}

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return regex.MatchString(email)
}
