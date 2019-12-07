package controllers

import (
	"../models"
	"../structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUsers(c *gin.Context){
	nama := c.PostForm("nama")
	username := c.PostForm("username")
	password := c.PostForm("password")
	tgl_lahir := c.PostForm("tgl_lahir")
	no_ktp := c.PostForm("no_ktp")
	no_hp := c.PostForm("no_hp")
	no_visa := c.PostForm("no_ktp")
	no_passpor := c.PostForm("no_passpor")
	foto, header, _ := c.Request.FormFile("foto")
	id_privileges := c.PostForm("id_privileges")

	response := structs.JsonResponse{}

	if nama == "" {
		response.ApiMessage = "isi nama"
	} else if username == "" {
		response.ApiMessage = "isi username"
	} else if password == "" {
		response.ApiMessage = "isi password"
	} else if tgl_lahir == "" {
		response.ApiMessage = "input tgl lahir"
	} else if no_ktp == "" {
		response.ApiMessage = "masukan nomor kependudukan anda"
	} else if no_visa == "" {
		response.ApiMessage = "masukan nomor visa anda"
	} else if no_passpor == "" {
		response.ApiMessage = "masukan nomor passpor anda"
	} else if no_hp == "" {
		response.ApiMessage = "masukan nomor hp anda"
	} else if id_privileges == "" {
		response.ApiMessage = "isi privileges"
	}

	response = models.CreateUsers(nama, username, password, tgl_lahir, no_ktp, no_hp, no_visa, no_passpor, foto, header, id_privileges)

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func LoginAdmin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	response := structs.JsonResponse{}

	if username == "" {
		response.ApiMessage = "Required username"
	} else if password == "" {
		response.ApiMessage = "Required Password"
	} else {
		response = models.LoginAdmin(username, password)
	}

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func LoginPetugas(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	response := structs.JsonResponse{}

	if username == "" {
		response.ApiMessage = "Required username"
	} else if password == "" {
		response.ApiMessage = "Required Password"
	} else {
		response = models.LoginPetugas(username, password)
	}

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func LoginUsers(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	response := structs.JsonResponse{}

	if username == "" {
		response.ApiMessage = "Required username"
	} else if password == "" {
		response.ApiMessage = "Required Password"
	} else {
		response = models.LoginUsers(username, password)
	}

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func GetUsers(c *gin.Context) {
	nama := c.Query("nama")
	username := c.Query("username")
	limit := c.Query("limit")
	offset := c.Query("offset")
	tgl_lahir := c.Query("tgl_lahir")
	no_ktp := c.Query("no_ktp")
	no_hp := c.Query("no_hp")
	no_visa := c.Query("no_ktp")
	no_passpor := c.Query("no_passpor")
	id_privileges := c.Query("id_privileges")

	response := structs.JsonResponse{}

	response = models.GetUsers(nama, username, tgl_lahir, no_ktp, no_hp, no_visa, no_passpor, id_privileges, offset, limit)

	if response.ApiStatus == 1 {

		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(500, response)
	}
}

func GetUserDetail(c *gin.Context) {
	id := c.Query("id")

	responses := structs.JsonResponse{}

	if id == "" {
		responses.ApiStatus = 0
		responses.ApiMessage = "Required Id"
		responses.Data = nil

	} else {
		responses = models.GetUserDetail(id)

	}
	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(400, responses)
	}
}

func UpdateUser(c *gin.Context)  {

	id := c.PostForm("id")
	nama := c.PostForm("nama")
	username := c.PostForm("username")
	password := c.PostForm("password")
	no_hp := c.PostForm("no_hp")
	no_visa := c.PostForm("no_visa")
	no_passpor := c.PostForm("no_passpor")
	foto, header, _ := c.Request.FormFile("foto")

	responses := structs.JsonResponse{}

	if id == "" {
		responses.ApiMessage = "Required Id"
	} else {

		responses = models.UpdateUser(id, nama, username, password, no_hp, no_visa, no_passpor, foto, header)
	}

	if responses.ApiStatus == 1 {

		c.JSON(http.StatusOK, responses)
	} else {
		c.JSON(500, responses)
	}

}