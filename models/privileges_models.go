package models

import (
	"../structs"
)

func GetRole()structs.JsonResponse {
	var (
		roles structs.GetPrivileges
		t structs.Component
	)
	response := structs.JsonResponse{}

	err := idb.DB.Table("privileges")

	err = err.Find(&roles)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = roles
	}

	return response
}

func GetType()structs.JsonResponse {
	var (
		getType structs.GetType
		t structs.Component
	)

	response := structs.JsonResponse{}
	err := idb.DB.Table("typerekom")

	err = err.Find(&getType)
	errx := err.Error

	if errx != nil {

		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil

	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = getType
	}

	return response
}
