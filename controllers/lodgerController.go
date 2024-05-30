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

func GetAllLodgers(c *gin.Context) {
	var result gin.H
	lodgers, err := repositories.GetAllLodgers(database.DBConnection)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"success": true,
			"message": "Berhasil mengambil seluruh data lodger",
			"data":    lodgers,
		}
		c.JSON(http.StatusOK, result)
	}

}

func InsertLodger(c *gin.Context) {
	var lodger structs.Lodger

	err := c.ShouldBindJSON(&lodger)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}

	err = repositories.InsertLodger(database.DBConnection, lodger)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})

	} else {

		c.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Berhasil menambahkan data lodger",
			"data":    utils.Empty,
		})
	}

}

func UpdateLodger(c *gin.Context) {

	var lodger structs.Lodger

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}
	err = c.ShouldBindJSON(&lodger)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}

	lodger.ID = id
	err = repositories.UpdateLodger(database.DBConnection, lodger)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Berhasil memperbarui data lodger",
			"data":    utils.Empty,
		})
	}

}

func DeleteLodger(c *gin.Context) {
	var lodger structs.Lodger
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}
	lodger.ID = id
	err = repositories.DeleteLodger(database.DBConnection, lodger)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Berhasil menghapus data lodger",
			"data":    utils.Empty,
		})
	}
}
