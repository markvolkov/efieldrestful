package resources

import (
	"efieldrestful/db"
	"efieldrestful/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetLeaderBoard(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		level := params["level"]
		track := params["track"]
		limit, err := strconv.Atoi(params["limit"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		isGlobal, err := strconv.ParseBool(params["global"])

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		leaderBoard := services.GetLeaderBoard(service, level, track, limit, isGlobal)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(leaderBoard)
	}
}
