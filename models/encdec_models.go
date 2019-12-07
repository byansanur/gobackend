package models

import (
	"../encrypt_tools"
)

var keys = []byte("warnaPassGo02019")

func DecryptPassword(password string) (string, error) {

	passDecrypt, err := encrypt_tools.Decrypt(keys, password)

	return string(passDecrypt), err

}

func EncryptPassword(password string) (string, error) {

	passEncrypt, err := encrypt_tools.Encrypt(keys, password)

	return passEncrypt, err
}
