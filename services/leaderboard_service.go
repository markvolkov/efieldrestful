package services

import (
	"efieldrestful/db"
	"efieldrestful/models"
)

func GetGlobalLeaderBoard(service db.DatabaseService, level string, track string, limit int) models.LeaderBoard{

}


func GetClassLeaderBoard(service db.DatabaseService, level string, track string, limit int, deviceId string) models.LeaderBoard{

}
