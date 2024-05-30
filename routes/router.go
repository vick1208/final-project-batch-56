package routes

import (
	"project-indekost/controllers"

	"github.com/gin-gonic/gin"
)

func MainServer() *gin.Engine {

	router := gin.Default()

	// route Penghuni Kost
	router.GET("/lodgers", controllers.GetAllLodgers)
	router.POST("/lodgers", controllers.BasicAuthMiddleware(), controllers.InsertLodger)
	router.PUT("/lodgers/:id", controllers.BasicAuthMiddleware(), controllers.UpdateLodger)
	router.DELETE("/lodgers/:id", controllers.BasicAuthMiddleware(), controllers.DeleteLodger)

	// route Kamar Kost
	router.GET("/rooms", controllers.GetAllRooms)
	router.POST("/rooms", controllers.BasicAuthMiddleware(), controllers.InsertRoom)
	router.PUT("/rooms/:id", controllers.BasicAuthMiddleware(), controllers.UpdateRoom)
	router.DELETE("/rooms/:id", controllers.BasicAuthMiddleware(), controllers.DeleteRoom)

	return router
}
