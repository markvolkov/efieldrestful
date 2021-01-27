package resources

import (
	"efieldrestful/db"
	"net/http"
)

func GetAllInstructors(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

	}
}

func GetInstructorsByInstitution(service db.DatabaseService)  http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

	}
}

func GetInstructorFromId(service db.DatabaseService)  http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteInstructorById(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

	}
}

