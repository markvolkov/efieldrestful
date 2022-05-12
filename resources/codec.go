package resources

import (
	"efieldrestful/models"
	"encoding/json"
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

func decodeStudentName(r *http.Request) *models.UpdateDeviceNamePayload {
	var toDecode models.UpdateDeviceNamePayload
	json.NewDecoder(r.Body).Decode(&toDecode)
	return &toDecode
}

func decodeDevice(r *http.Request) *models.Device {
	var toDecode models.Device
	json.NewDecoder(r.Body).Decode(&toDecode)
	return &toDecode
}

func encodeLeaderboard(leaderBoard *models.LeaderBoard, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(leaderBoard)
}


func encodeError(w http.ResponseWriter, error string, status int) {
	http.Error(w, error, status)
}


