package routes

import (
	"project-indekost/controllers"

	"github.com/gin-gonic/gin"
)

func MainServer() *gin.Engine {

	router := gin.Default()

	// route Penghuni Kost
	router.GET("/lodgers", controllers.GetAllLodgers)
	router.POST("/lodgers", controllers.InsertLodger)
	router.PUT("/lodgers/:id", controllers.UpdateLodger)
	router.DELETE("/lodgers/:id", controllers.DeleteLodger)

	// route Kamar Kost
	router.GET("/rooms", controllers.GetAllRooms)
	router.POST("/rooms", controllers.InsertRoom)
	router.PUT("/rooms/:id", controllers.UpdateRoom)
	router.DELETE("/rooms/:id", controllers.DeleteRoom)

	return router
}
