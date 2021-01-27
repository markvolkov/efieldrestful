package services

import (
	"efieldrestful/db"
	"efieldrestful/models"
)

func GetLeaderBoard(service db.DatabaseService, level string, track string, limit int, isGlobal bool) models.LeaderBoard{
	if isGlobal {

	} else {

	}
}
