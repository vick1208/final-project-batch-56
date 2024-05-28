package controllers

import (
	"fmt"
	"project-indekost/database"
	"project-indekost/repositories"

	"github.com/gin-gonic/gin"
)

func GetAllLodgers(c *gin.Context) {

	lodgers, err := repositories.GetAllLodgers(database.DBConnection)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(lodgers)

}

func InsertLodger(c *gin.Context) {

}
