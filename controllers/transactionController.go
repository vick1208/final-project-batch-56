package controllers

import (
	"errors"
	"net/http"
	"project-indekost/database"
	"project-indekost/repositories"
	"project-indekost/structs"
	"project-indekost/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetLodgerDueDateTransaction(ctx *gin.Context) {
	var result gin.H
	transactions, err := repositories.GetDueDateTransaction(database.DBConnection)
	if err != nil {
		result = gin.H{
			"success": false,
			"message": errors.New("data gagal diambil").Error(),
			"data":    utils.Empty,
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, result)
	} else {
		result = gin.H{
			"success": true,
			"message": "Berhasil mengambil seluruh data transaction dengan id penghuni kost",
			"data":    transactions,
		}
		ctx.JSON(http.StatusOK, result)
	}
}

func InsertDataTransaction(ctx *gin.Context) {
	var transaction structs.Transaction
	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}
	err = repositories.InsertTransactionRentData(database.DBConnection, transaction)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": errors.New("data gagal ditambahkan").Error(),
			"data":    utils.Empty,
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"success": true,
			"message": "Berhasil menambahkan data transaction",
			"data":    utils.Empty,
		})
	}
}

func DeleteDataTransaction(ctx *gin.Context) {
	var transaction structs.Transaction
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    utils.Empty,
		})
	}

	transaction.ID = id
	err = repositories.DeleteTransactionRentData(database.DBConnection, transaction)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": errors.New("data gagal dihapus").Error(),
			"data":    utils.Empty,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Berhasil menghapus data transaction",
			"data":    utils.Empty,
		})
	}
}
