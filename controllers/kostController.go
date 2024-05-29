package controllers

import (
	"net/http"
	"project-indekost/database"
	"project-indekost/repositories"
	"project-indekost/structs"
	"strconv"

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

func UpdateLodger(c *gin.Context) {

	var lodger structs.Lodger

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	err = c.ShouldBindJSON(&lodger)
	if err != nil {
		panic(err)
	}

	lodger.ID = id
	err = repositories.UpdateLodger(database.DBConnection, lodger)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Lodger Data",
	})

}

func DeleteLodger(c *gin.Context) {
	var lodger structs.Lodger
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	lodger.ID = id
	err = repositories.DeleteLodger(database.DBConnection, lodger)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Lodger Data",
	})
}

func GetAllRooms(c *gin.Context) {
	var result gin.H
	rooms, err := repositories.GetAllRooms(database.DBConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	}
	result = gin.H{
		"result": rooms,
	}
	c.JSON(http.StatusOK, result)

}

func InsertRoom(c *gin.Context) {
	var room structs.Room

	err := c.ShouldBindJSON(&room)
	if err != nil {
		panic(err)
	}

	err = repositories.InsertRoom(database.DBConnection, room)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusCreated, gin.H{
		"result": "Success Insert Room Data",
	})
}

func UpdateRoom(c *gin.Context) {

	var room structs.Room

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	err = c.ShouldBindJSON(&room)
	if err != nil {
		panic(err)
	}
	room.ID = id
	err = repositories.UpdateRoom(database.DBConnection, room)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Lodger Data",
	})

}
