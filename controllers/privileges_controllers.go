package controllers

import (
	"github.com/gin-gonic/gin"
	"../structs"
	"../models"
	"net/http"
)

func GetRole(c *gin.Context) {
	responses := structs.JsonResponse{}
	responses = models.GetRole()

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}
}

func GetType(c *gin.Context){
	responses := structs.JsonResponse{}
	responses = models.GetType()

	if responses.ApiStatus == 1 {
		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}
