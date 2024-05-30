package controllers

import (
	"net/http"
	"project-indekost/database"
	"project-indekost/repositories"
	"project-indekost/utils"

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
