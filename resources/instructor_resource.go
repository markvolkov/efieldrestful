package resources

import (
	"efieldrestful/db"
	"efieldrestful/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllInstructors(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		mongoResult := services.GetAllInstructors(service)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mongoResult)
	}
}

func GetInstructorsByInstitution(service db.DatabaseService)  http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		mongoResult := services.GetInstructorsByInstitution(service, params["institution"])
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mongoResult)
	}
}

func GetInstructorFromId(service db.DatabaseService)  http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		mongoResult := services.GetInstructorFromId(service, params["instructor_id"])
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mongoResult)
	}
}

func DeleteInstructorById(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		mongoResult := services.DeleteInstructorById(service, params["instructor_id"])
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mongoResult)
	}
}

