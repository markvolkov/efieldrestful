package encrypt

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

const cost = 12

func EncryptData(data []byte) []byte {
	result, err := bcrypt.GenerateFromPassword(data, cost)
	checkError(err)
	return result
}

func CheckData(hashedData []byte, data []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedData, data)
	checkError(err)
	return err == nil
}

func checkError(err error) {
	if err != nil {
		log.Println("Error: " + err.Error())
	}
}
