package resources

import (
	"efieldrestful/db"
	"net/http"
)

func GetLeaderBoard(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		//params := mux.Vars(r)
		//level := params["level"]
		//track := params["track"]
		//limit := params["limit"]

	}
}
