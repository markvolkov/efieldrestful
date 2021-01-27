package resources

import (
	"efieldrestful/db"
	"efieldrestful/services"
	"encoding/json"
	"net/http"
	"strconv"
)

//leaderboard handlers usage GET /leaderboard?limit=10&level=1&track=1&global=false/
func GetLeaderBoard(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		level := r.URL.Query().Get("level")
		track := r.URL.Query().Get("track")
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		isGlobal, err := strconv.ParseBool(r.URL.Query().Get("global"))

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
