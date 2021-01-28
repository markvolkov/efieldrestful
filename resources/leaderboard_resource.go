package resources

import (
	"efieldrestful/db"
	"efieldrestful/services"
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

		if isGlobal {
			leaderBoard := services.GetGlobalLeaderBoard(service, level, track, limit)
			encodeLeaderboard(leaderBoard, w)
		} else {
			deviceId := r.URL.Query().Get("device_id")
			leaderBoard := services.GetClassLeaderBoard(service, level, track, limit, deviceId)
			encodeLeaderboard(leaderBoard, w)
		}
	}
}
