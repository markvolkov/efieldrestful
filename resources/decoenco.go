package resources

import (
	"efieldrestful/models"
	"encoding/json"
	"log"
	"net/http"
)

/**
    ========================
	Json Encoders / Decoders
    ========================
 */

func decodeAttempt(r *http.Request) *models.Attempt {
	var toDecode models.Attempt
	json.NewDecoder(r.Body).Decode(&toDecode)
	return &toDecode
}

func decodeClass(r *http.Request) *models.Class {
	var toDecode models.Class
	json.NewDecoder(r.Body).Decode(&toDecode)
	return &toDecode
}

func encodeError(w http.ResponseWriter, error string, status int) {
	http.Error(w, error, status)
}


func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}

