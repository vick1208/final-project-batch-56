package controllers

import (
	"net/http"
	"project-indekost/utils"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware() gin.HandlerFunc {

	validUsers := gin.Accounts{
		"owner": "owner",
	}

	return func(ctx *gin.Context) {
		username, password, ok := ctx.Request.BasicAuth()
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Username dan password perlu dimasukkan",
				"data":    utils.Empty,
			})

			ctx.Abort()
			return
		}

		if validPass, found := validUsers[username]; !found || validPass != password {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Username atau password salah",
				"data":    utils.Empty,
			})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
