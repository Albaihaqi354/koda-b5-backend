package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	app.POST("/auth/register", func(ctx *gin.Context) {
		var register User
		if err := ctx.ShouldBindJSON(&register); err != nil {
			log.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "Error",
			})
			return
		}

		IsValid(register.Email)

		ctx.JSON(http.StatusOK, gin.H{
			"data": register,
		})
	})

	app.Run("localhost:5000")
}

func IsValid(v string) {
	matched := regexp.MustCompile(`!/^[^\s@]+@[^\s@]+\.[^\s@]+$/`)
	return matched.MatchString(v)
}

type User struct {
	Email    string `form:"Email"`
	Password string `form:"password"`
}
