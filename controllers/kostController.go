package controllers

import (
	"net/http"
	"project-indekost/database"
	"project-indekost/repositories"

	"github.com/gin-gonic/gin"
)

func GetAllLodgers(c *gin.Context) {
	var result gin.H
	lodgers, err := repositories.GetAllLodgers(database.DBConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	}
	result = gin.H{
		"result": lodgers,
	}
	c.JSON(http.StatusOK, result)
}
