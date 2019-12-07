package controllers

import (
	"github.com/gin-gonic/gin"
	"../structs"
	"../models"
	"net/http"
)

func CreateRekomendasi(c *gin.Context) {
	nama := c.PostForm("nama")
	alamat := c.PostForm("alamat")
	foto, header, _ := c.Request.FormFile("foto")
	lat := c.PostForm("lat")
	lng := c.PostForm("lng")
	rating := c.PostForm("rating")
	id_type := c.PostForm("id_type")

	response := structs.JsonResponse{}

	if nama == "" {
		response.ApiMessage = "isi title"
	} else if alamat == "" {
		response.ApiMessage = "isi alamat"
	} else if lat == "" {
		response.ApiMessage = "latitude"
	} else if lng == "" {
		response.ApiMessage = "longitude"
	} else if rating == "" {
		response.ApiMessage = "rating required"
	} else if id_type == "" {
		response.ApiMessage = "type isi"
	}

	response = models.CreateRekomendasi(nama, alamat, foto, header, lat, lng, rating, id_type)

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}

}
func GetRekom(c *gin.Context){
	nama := c.Query("nama")
	alamat := c.Query("alamat")
	foto := c.Query("foto")
	rating := c.Query("rating")
	id_type := c.Query("id_type")
	offset := c.Query("offset")
	limit := c.Query("limit")

	res := structs.JsonResponse{}

	res = models.GetRekomends(nama, alamat, foto, rating, id_type, offset,limit)

	if res.ApiStatus == 1 {

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(500, res)
	}
}
func GetRekomDetail(c *gin.Context) {
	id := c.Query("id")

	res := structs.JsonResponse{}

	if id == "" {
		res.ApiStatus = 0
		res.ApiMessage = "Required id"
		res.Data = nil
	}else {
		res = models.GetRekomendsDetail(id)

	}
	if res.ApiStatus == 1 {

		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(400, res)
	}

}

func UpdateRekom(c *gin.Context) {
	id := c.PostForm("id")
	nama := c.PostForm("nama")
	alamat := c.PostForm("alamat")
	foto, header, _ := c.Request.FormFile("foto")
	rating := c.PostForm("rating")

	responses := structs.JsonResponse{}

	if id == "" {
		responses.ApiMessage = "Required Id"
	} else {
		responses = models.UpdateRekomendasi(id, nama, alamat, foto, header, rating)
	}
	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}
}
