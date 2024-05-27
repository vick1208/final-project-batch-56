package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainServer() *gin.Engine {

	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK")
	})
	return router
}
