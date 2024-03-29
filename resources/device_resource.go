package resources

import (
	"efieldrestful/db"
	"efieldrestful/models"
	"efieldrestful/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)


func StoreAttempt(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		if r.Body == http.NoBody {
			mongoResult := &models.Device{DeviceId: primitive.NewObjectID(), Attempts: make([]models.Attempt, 0)}
			result := services.StoreDevice(service, mongoResult)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&result)
		} else {
			attemptDecoded := decodeAttempt(r)
			mongoResult := services.GetDeviceById(service, attemptDecoded.AttemptId.Hex())
			if mongoResult == nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			attemptDecoded.AttemptId = primitive.NewObjectID()
			mongoResult.Attempts = append(mongoResult.Attempts, *attemptDecoded)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			result := services.StoreAttemptFromDevice(service, mongoResult.DeviceId.Hex(), *attemptDecoded)
			json.NewEncoder(w).Encode(&result)
		}
	}
}

func GetAttemptsByDeviceId(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		deviceId := params["deviceId"]
		log.Println(deviceId)
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

func UpdateStudentName(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		deviceId := params["deviceId"]
		mongoResult := services.GetDeviceById(service, deviceId)
		if mongoResult == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			encodeError(w, "A Device With That ID Doesn't Exists.", http.StatusBadRequest)
		} else {
			//update student name after decoding name payload
			deviceUpdate := decodeStudentName(r)
			updateResponse := services.UpdateDeviceName(service, mongoResult.DeviceId.Hex(), deviceUpdate.StudentName)
			log.Println("Updating student-name for device-id: " + deviceId + " to " + deviceUpdate.StudentName)
			updateResponse = services.UpdateDeviceClassName(service, mongoResult.DeviceId.Hex(), deviceUpdate.ClassName)
			log.Println("Updating class-name for device-id: " + deviceId + " to " + deviceUpdate.ClassName)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updateResponse)
		}
	}
}

func GetDevicesByStudentName(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		mongoResult := services.GetDevicesByStudentName(service, params["student_name"])
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mongoResult)
	}
}


func DeleteDeviceById(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err, mongoResult := services.DeleteDeviceById(service, params["device_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mongoResult)
		}
	}
}

