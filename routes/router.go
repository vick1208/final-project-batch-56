package routes

import (
	"project-indekost/controllers"

	"github.com/gin-gonic/gin"
)

func MainServer() *gin.Engine {

	r := gin.Default()

	r.GET("/lodgers", controllers.GetAllLodgers)
	return r
}
