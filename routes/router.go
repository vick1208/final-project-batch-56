package routes

import (
	"project-indekost/controllers"

	"github.com/gin-gonic/gin"
)

func MainServer() *gin.Engine {

	router := gin.Default()

	router.GET("/lodgers", controllers.GetAllLodgers)
	router.POST("/lodgers", controllers.InsertLodger)
	router.PUT("/lodgers/:id", controllers.UpdateLodger)
	router.DELETE("/lodgers/:id", controllers.DeleteLodger)

	return router
}
