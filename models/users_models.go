package models

import (
	"../structs"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"mime/multipart"
	"strconv"
	"time"
	"../config_db"
)


// fungsi unutk membuat user baru
// parameter di inisialisasi diatas sesuai pada structnya
func CreateUsers(nama string, username string, password string,
	tgl_lahir string, no_ktp string, no_hp string, no_visa string,
	no_passpor string, files multipart.File,
	header *multipart.FileHeader, id_privileges string) structs.JsonResponse  {

	// pembuatan variable dengan mengarahkan ke struct yang dibuat
	var (
		users 	structs.CreateUsers
		CheckId structs.CheckId
		t 		structs.Component
	)

	// pemanggilan json response dengan membuat variable dan panggil struct
	response := structs.JsonResponse{}

	// query mysql
	check := idb.DB.Table("users").Select("count(users.id) as count_id")

	// pemberitahuan pada db untuk variable username ada pada table users dan field username
	check = check.Where("users.username = ?", username)

	// check first id
	check = check.First(&CheckId)

	// untuk cek error
	checkx := check.Error

	// jika error itu kosong atau nil, maka tidak ada error
	if checkx == nil {
		fmt.Println("cek id ga error")

		// cek id jika id = 0 maka password di encrypt pada database dan hasil encrypt bukan hasil convert
		// melainkan hilang pada database
		if CheckId.CountId == 0 {
			encrptPassword, _ := EncryptPassword(password)
			id_privileges_conv, _ := strconv.Atoi(id_privileges)

			url := UploadImage("user", fmt.Sprint(username), files, header)

			if url != "" {
				fmt.Println("foto ga null")

				// pengenalan variable setiap field
				users.Nama = nama
				users.Username = username
				users.Password = encrptPassword
				users.TglLahir = tgl_lahir
				users.NoKtp = no_ktp
				users.NoHp = no_hp
				users.NoVisa =no_visa
				users.NoPasspor = no_passpor
				users.Foto = url
				users.IdPrivileges = id_privileges_conv
				users.CreatedAt = t.GetTimeNow()

				// query untuk insert ke table users
				err := idb.DB.Table("users").Create(&users)

				// jika gagal insert ke table
				errx := err.Error

				// kondisi jika gagal buat akun
				if errx != nil {
					fmt.Println("gagal buat akun")
					response.ApiMessage = t.GetMessageErr()
				} else {
					// jika berhasil buat akun atau insert ke db maka responsenya disini
					response.ApiStatus = 1
					response.ApiMessage = t.GetMessageSucc()
					response.Data = users
				}
			}
		} else {
			// jika user mendaftar dengan username yang sama maka akan response disini
			response.ApiMessage = "Username Already Used"
		}
	} else {
		// jika gagal cek id
		fmt.Println(checkx)
	}
	// kembalian response
	return  response
}

func LoginAdmin(username string, password string) structs.JsonResponse {

	var (
		userlogin structs.CekUserLogin
		users structs.GetUserLogin
	)

	response := structs.JsonResponse{}

	cekUsername := idb.DB.Table("users").Select("id,username,password")
	cekUsername = cekUsername.Where("users.username = ?", username)

	cekUsername = cekUsername.Scan(&userlogin)
	cekUsernames := cekUsername.RecordNotFound()

	fmt.Println(userlogin.Id)

	if cekUsernames {
		fmt.Println("belum terdaftar")
		response.ApiMessage = "belum terdaftar"
	} else {
		cekPassword, errPass := DecryptPassword(userlogin.Password)

		if errPass != nil {
			fmt.Println("username/password salah")
			response.ApiMessage = "Password salah"
		} else {
			fmt.Println("ok ada nih")
			if cekPassword == password {
				fmt.Println("Pass sama")
				err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, privileges.role, users.created_at")
				err = err.Joins("join privileges on users.id_privileges = privileges.id")
				err = err.Where("users.id = ?", userlogin.Id)
				err = err.Where("users.id_privileges != 2")
				err = err.Where("users.id_privileges != 3")
				err = err.First(&users)
				errx := err.Error

				sign := jwt.New(jwt.SigningMethodHS256)
				claim := sign.Claims.(jwt.MapClaims)

				claim["authorized"] = true
				claim["client"] = users.Id
				claim["exp"] = time.Now().Add(time.Minute * 360).Unix()

				token, _ := sign.SignedString(config_db.JwtKey())
				users.Token = token

				if errx == nil {
					response.ApiStatus = 1
					response.ApiMessage = "success login"
					response.Data = users
				} else {
					response.ApiMessage = "gagal login"
				}
			} else {
				response.ApiMessage = "password salah"
			}
		}
	}
	return response
}

func LoginPetugas(username string, password string) structs.JsonResponse {

	var (
		userlogin structs.CekUserLogin
		users structs.GetUserLogin
	)

	response := structs.JsonResponse{}

	cekUsername := idb.DB.Table("users").Select("id,username,password")
	cekUsername = cekUsername.Where("users.username = ?", username)

	cekUsername = cekUsername.Scan(&userlogin)
	cekUsernames := cekUsername.RecordNotFound()

	fmt.Println(userlogin.Id)

	if cekUsernames {
		fmt.Println("belum terdaftar")
		response.ApiMessage = "belum terdaftar"
	} else {
		cekPassword, errPass := DecryptPassword(userlogin.Password)

		if errPass != nil {
			fmt.Println("username/password salah")
			response.ApiMessage = "Password salah"
		} else {
			fmt.Println("ok ada nih")
			if cekPassword == password {
				fmt.Println("Pass sama")
				err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, privileges.role, users.created_at")
				err = err.Joins("join privileges on users.id_privileges = privileges.id")
				err = err.Where("users.id = ?", userlogin.Id)
				err = err.Where("users.id_privileges != 1")
				err = err.Where("users.id_privileges != 3")
				err = err.First(&users)
				errx := err.Error

				sign := jwt.New(jwt.SigningMethodHS256)
				claim := sign.Claims.(jwt.MapClaims)

				claim["authorized"] = true
				claim["client"] = users.Id
				claim["exp"] = time.Now().Add(time.Minute * 360).Unix()

				token, _ := sign.SignedString(config_db.JwtKey())
				users.Token = token

				if errx == nil {
					response.ApiStatus = 1
					response.ApiMessage = "success login"
					response.Data = users
				} else {
					response.ApiMessage = "gagal login"
				}
			} else {
				response.ApiMessage = "password salah"
			}
		}
	}
	return response
}

func LoginUsers(username string, password string) structs.JsonResponse {

	var (
		userlogin structs.CekUserLogin
		users structs.GetUserLogin
	)

	response := structs.JsonResponse{}

	cekUsername := idb.DB.Table("users").Select("id,username,password")
	cekUsername = cekUsername.Where("users.username = ?", username)

	cekUsername = cekUsername.Scan(&userlogin)
	cekUsernames := cekUsername.RecordNotFound()

	fmt.Println(userlogin.Id)

	if cekUsernames {
		fmt.Println("belum terdaftar")
		response.ApiMessage = "belum terdaftar"
	} else {
		cekPassword, errPass := DecryptPassword(userlogin.Password)

		if errPass != nil {
			fmt.Println("username/password salah")
			response.ApiMessage = "Password salah"
		} else {
			fmt.Println("ok ada nih")
			if cekPassword == password {
				fmt.Println("Pass sama")
				err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, privileges.role, users.created_at")
				err = err.Joins("join privileges on users.id_privileges = privileges.id")
				err = err.Where("users.id = ?", userlogin.Id)
				err = err.Where("users.id_privileges != 1")
				err = err.Where("users.id_privileges != 2")
				err = err.First(&users)
				errx := err.Error

				sign := jwt.New(jwt.SigningMethodHS256)
				claim := sign.Claims.(jwt.MapClaims)

				claim["authorized"] = true
				claim["client"] = users.Id
				claim["exp"] = time.Now().Add(time.Minute * 360).Unix()

				token, _ := sign.SignedString(config_db.JwtKey())
				users.Token = token

				if errx == nil {
					response.ApiStatus = 1
					response.ApiMessage = "success login"
					response.Data = users
				} else {
					response.ApiMessage = "gagal login"
				}
			} else {
				response.ApiMessage = "password salah"
			}
		}
	}
	return response
}

func GetUsers(nama string, username string, tgl_lahir string, no_ktp string,
	no_hp string, no_visa string, no_passpor string, id_privileges string, offset string, limit string)structs.JsonResponse {
	var (
		getUser []structs.GetUser
		t structs.Component
	)

	response := structs.JsonResponse{}

	err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, users.created_at," + "privileges.privileges")
	err = err.Joins("JOIN privileges ON users.id_privileges = privileges.id")

	if limit != "" {
		err = err.Limit(limit)
	}
	if offset != "" {
		err = err.Offset(offset)
	}
	if nama != "" {
		err = err.Where("users.nama = ?", nama)
	}
	if username != "" {
		err = err.Where("users.username = ?", username)
	}
	if tgl_lahir != "" {
		err = err.Where("users.tgl_lahir = ?", tgl_lahir)
	}
	if no_ktp != "" {
		err = err.Where("users.no_ktp = ?", no_ktp)
	}
	if no_hp != "" {
		err = err.Where("users.no_hp = ?", no_hp)
	}
	if no_visa != "" {
		err = err.Where("users.no_visa = ?", no_visa)
	}
	if no_passpor != "" {
		err = err.Where("users.no_passpor = ?", no_passpor)
	}
	if id_privileges != "" {
		err = err.Where("users.id_privileges = ?", id_privileges)
	}

	err = err.Order("users.nama asc")

	err = err.Find(&getUser)
	errx := err.Error

	if errx != nil{
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = getUser
	}
	return response
}

func GetUserDetail(id string)structs.JsonResponse {

	var (
		users structs.GetUser
		t structs.Component
	)

	response := structs.JsonResponse{}
	err := idb.DB.Table("users").Select("users.id, users.nama, users.username, users.tgl_lahir, users.no_ktp, users.no_hp, users.no_visa, users.no_passpor, users.foto, users.id_privileges, users.created_at," + "privileges.privileges")
	err = err.Joins("JOIN privileges ON users.id_privileges = privileges.id")
	err = err.Where("users.id = ?", id)
	err = err.First(&users)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = users
	}
	return response
}

func UpdateUser(id string, nama string, username string, password string,
	no_hp string, no_visa string, no_passpor string, files multipart.File, header *multipart.FileHeader,) structs.JsonResponse {
	var (
		userUpdate structs.UpdateUsers
		t structs.Component
	)
	response := structs.JsonResponse{}
	encryptPassword, _ := EncryptPassword(password)
	id_conv, _ := strconv.Atoi(id)
	no_hp_conv, _ := strconv.Atoi(no_hp)

	url := UploadImage("user", fmt.Sprint(username), files, header)

	if password != "" {
		userUpdate.Password = encryptPassword
	}

	userUpdate.Nama = nama
	userUpdate.Username = username
	userUpdate.NoHp = no_hp_conv
	userUpdate.NoVisa = no_visa
	userUpdate.NoPasspor = no_passpor
	userUpdate.Foto = url

	err := idb.DB.Table("users").Where("users.id = ?", id_conv).Update(&userUpdate)
	errx := err.Error

	if errx != nil {
		response.ApiStatus = 0
		response.ApiMessage = errx.Error()
		response.Data = nil
	} else {
		response.ApiStatus = 1
		response.ApiMessage = t.GetMessageSucc()
		response.Data = userUpdate
	}
	return response
}