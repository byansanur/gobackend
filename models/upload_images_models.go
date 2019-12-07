package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UploadImage(type_foto string, id_data string, files multipart.File, header *multipart.FileHeader) string {

	b := make([]rune, 40)
	bc := make([]rune, 16)
	t := time.Now()
	pathImage := ""
	pathId := []byte(id_data)
	pathIdMD5 := md5.Sum(pathId)
	var letterRunes = []rune("abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	filename := strings.Join(strings.Fields(header.Filename), "")

	for i := range b {

		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	for i := range bc {

		bc[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	filename = string(b) + t.Format("20060102150405.0") + string(bc) + filename
	folderID := hex.EncodeToString(pathIdMD5[:])

	pathImage = "./pbc/img/all/" + folderID + "/"

	errMk := os.MkdirAll(pathImage, 0777)

	if errMk != nil {
		fmt.Println(errMk.Error())
		return ""
	} else {

		out, err := os.Create(pathImage + filename)

		if err != nil {
			fmt.Println(err.Error())
			return ""
		} else {

			defer out.Close()

			_, err = io.Copy(out, files)

			if err != nil {
				fmt.Println(err.Error())
				return ""
			}

			// return string(filepath.Dir(pathImage)+filepath.FromSlash("/"+filename))
			return string(filepath.FromSlash(folderID + "/" + filename))

		}

	}

	//return
}

//func GetPhotoVisumTarget(id_target_visum string) structs.JsonResponse {
//
//	var (
//		targetvisumdetail []structs.TargetVisumPhoto
//		t                 structs.Component
//	)
//
//	response := structs.JsonResponse{}
//	//tables := t.TblLead()
//
//	id_target_visum_conv, _ := strconv.Atoi(id_target_visum)
//
//	err := idb.DB.Table("target_photo").Select(`target_photo.id,
//target_photo.id_target_visum, target_photo.photo`)
//
//	err = err.Where("target_photo.id_target_visum = ?", id_target_visum_conv)
//
//	err = err.Find(&targetvisumdetail)
//	errx := err.Error
//
//	if errx != nil {
//
//		response.ApiStatus = 0
//		response.ApiMessage = errx.Error()
//		response.Data = nil
//
//	} else {
//		response.ApiStatus = 1
//		response.ApiMessage = t.GetMessageSucc()
//		response.Data = targetvisumdetail
//	}
//
//	return response
//}
//
//func GetPhotoVisumLead(id_lead_visum string) structs.JsonResponse {
//
//	var (
//		targetvisumdetail []structs.LeadVisumPhoto
//		t                 structs.Component
//	)
//
//	response := structs.JsonResponse{}
//	//tables := t.TblLead()
//
//	id_lead_visum_conv, _ := strconv.Atoi(id_lead_visum)
//
//	err := idb.DB.Table("lead_photo").Select(`lead_photo.id,
//lead_photo.id_lead_visum, lead_photo.photo`)
//
//	err = err.Where("lead_photo.id_lead_visum = ?", id_lead_visum_conv)
//
//	err = err.Find(&targetvisumdetail)
//	errx := err.Error
//
//	if errx != nil {
//
//		response.ApiStatus = 0
//		response.ApiMessage = errx.Error()
//		response.Data = nil
//
//	} else {
//		response.ApiStatus = 1
//		response.ApiMessage = t.GetMessageSucc()
//		response.Data = targetvisumdetail
//	}
//
//	return response
//}
//
//func GetPhotoOrder(id_order string) structs.JsonResponse {
//
//	var (
//		targetvisumdetail []structs.GetOrderPhoto
//		t                 structs.Component
//	)
//
//	response := structs.JsonResponse{}
//	//tables := t.TblLead()
//
//	id_order_conv, _ := strconv.Atoi(id_order)
//
//	err := idb.DB.Table("order_photo").Select(`order_photo.id,
//order_photo.id_order, order_photo.photo`)
//
//	err = err.Where("order_photo.id_order = ?", id_order_conv)
//
//	err = err.Find(&targetvisumdetail)
//	errx := err.Error
//
//	if errx != nil {
//
//		response.ApiStatus = 0
//		response.ApiMessage = errx.Error()
//		response.Data = nil
//
//	} else {
//		response.ApiStatus = 1
//		response.ApiMessage = t.GetMessageSucc()
//		response.Data = targetvisumdetail
//	}
//
//	return response
//}
