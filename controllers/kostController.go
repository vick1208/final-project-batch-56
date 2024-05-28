package controllers

import (
	"net/http"
	"project-indekost/database"
	"project-indekost/repositories"
	"project-indekost/structs"

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

func InsertLodger(c *gin.Context) {
	var lodger structs.Lodger

	err := c.ShouldBindJSON(&lodger)
	if err != nil {
		panic(err)
	}

	err = repositories.InsertLodger(database.DBConnection, lodger)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": "Success Insert Lodger Data",
	})
}
