package controllers

import (
	"net/http"
	"project-indekost/database"
	"project-indekost/repositories"
	"project-indekost/structs"
	"project-indekost/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllRooms(c *gin.Context) {
	var result gin.H
	rooms, err := repositories.GetAllRooms(database.DBConnection)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		}
	} else {
		result = gin.H{
			"success": true,
			"message": "Berhasil mengambil seluruh data room",
			"data":    rooms,
		}
	}
	c.JSON(http.StatusOK, result)

}

func InsertRoom(c *gin.Context) {
	var room structs.Room

	err := c.ShouldBindJSON(&room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}

	err = repositories.InsertRoom(database.DBConnection, room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Berhasil menambahkan data room",
			"data":    utils.Empty,
		})
	}

}

func UpdateRoom(c *gin.Context) {

	var room structs.Room

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}
	err = c.ShouldBindJSON(&room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}
	room.ID = id
	err = repositories.UpdateRoom(database.DBConnection, room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Berhasil memperbarui data room",
			"data":    utils.Empty,
		})
	}

}

func DeleteRoom(c *gin.Context) {
	var room structs.Room

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}
	room.ID = id
	err = repositories.DeleteRoom(database.DBConnection, room)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Berhasil menghapus data room",
			"data":    utils.Empty,
		})
	}
}
