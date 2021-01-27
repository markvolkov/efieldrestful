package resources

import (
	"efieldrestful/db"
	"efieldrestful/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateClass(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		classPayload := decodeClass(r)
		mongoResult := services.GetClassFromId(service, classPayload.ClassId.Hex())
		if mongoResult == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			result := services.SaveClass(service, classPayload)
			json.NewEncoder(w).Encode(&result)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			encodeError(w, "A Class With That ID Already Exists.", http.StatusBadRequest)
		}
	}
}


func GetClass(service db.DatabaseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		mongoResult := services.GetClassFromId(service, params["classId"])
		if mongoResult == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			encodeError(w, "A Class With That ID Doesn't Exists.", http.StatusNotFound)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mongoResult)
		}
	}
}

func ClassList(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		mongoResult := services.GetAllClasses(service)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mongoResult)
	}
}


func GetDevicesFromClass(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

	}
}

func StoreDeviceInClass(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteClassByAccessCode(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteClassById(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

	}
}
