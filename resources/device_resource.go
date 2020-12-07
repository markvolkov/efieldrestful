package resources

import (
	"efieldrestful/db"
	"efieldrestful/models"
	"efieldrestful/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func StoreAttempt(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		deviceId := params["deviceId"]
		mongoResult := services.GetDeviceById(service, deviceId)
		if mongoResult == nil {
			mongoResult = &models.Device{DeviceId: primitive.NewObjectID(), Attempts: make([]models.Attempt, 0)}
		}
		attempt := decodeAttempt(r)
		mongoResult.Attempts = append(mongoResult.Attempts, *attempt)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		result := services.StoreAttemptFromDevice(service, deviceId, *attempt)
		json.NewEncoder(w).Encode(&result)
	}
}

func GetAttemptsByDeviceId(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		deviceId := params["deviceId"]
		mongoResult := services.GetDeviceById(service, deviceId)
		if mongoResult == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			encodeError(w, "A Device With That ID Doesn't Exists.", http.StatusBadRequest)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mongoResult)
		}
	}
}

func GetDevices(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		mongoResult := services.GetAllDevices(service)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mongoResult)
	}
}

