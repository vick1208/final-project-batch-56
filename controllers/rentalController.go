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

func GetAllDataRental(ctx *gin.Context) {
	var result gin.H
	rents, err := repositories.GetDataRental(database.DBConnection)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"success": true,
			"message": "Berhasil mengambil seluruh data rental",
			"data":    rents,
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func InsertDataRental(ctx *gin.Context) {
	var rental structs.Rental
	err := ctx.ShouldBindJSON(&rental)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}
	err = repositories.InsertDataRental(database.DBConnection, rental)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Berhasil menambahkan data rental",
			"data":    utils.Empty,
		})
	}
}

func DeleteDataRental(ctx *gin.Context) {
	var rental structs.Rental
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}
	rental.ID = id
	err = repositories.DeleteDataRental(database.DBConnection, rental)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Berhasil menghapus data rental",
			"data":    utils.Empty,
		})
	}
}
