package routes

import (
	"project-indekost/controllers"

	"github.com/gin-gonic/gin"
)

func MainServer() *gin.Engine {

	r := gin.Default()
	owner := r.Group("/owner")
	owner.GET("/lodgers", controllers.GetAllLodgers)
	return r
}
