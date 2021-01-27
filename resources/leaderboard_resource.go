package resources

import (
	"efieldrestful/db"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetLeaderBoard(service db.DatabaseService) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		level := params["level"]
		track := params["track"]
		limit := params["limit"]
		isGlobal, err := strconv.ParseBool(params["global"])

		if err != nil || !isGlobal  {
			//Either there was an error parsing the boolean which defaults to not wanting global or you don't want the global leaderboard
		} else {

		}

	}
}
