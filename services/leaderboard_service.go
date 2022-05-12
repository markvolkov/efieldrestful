package services

import (
	"efieldrestful/db"
	"efieldrestful/models"
	"log"
	"sort"
	"strings"
)

func GetGlobalLeaderBoard(service db.DatabaseService, level string, limit int) *models.LeaderBoard {
	boardAttempts := make([]models.LeaderBoardStat, 0)
	devices := GetAllDevices(service)
	for _, device := range devices {
		bestAttempt := models.Attempt{Score: 0}
		for _, attempt := range device.Attempts {
			log.Println(attempt.AttemptId.String(), attempt.Score, attempt.Level, level)
			if strings.Compare(attempt.Level, level) != 0 {
				continue
			}
			if bestAttempt.Score == 0 {
				bestAttempt = attempt
			} else {
				if attempt.Score > bestAttempt.Score {
					bestAttempt = attempt
				}
			}
		}
		if bestAttempt.Score != 0 {
			boardAttempts = append(boardAttempts, models.LeaderBoardStat{TopAttempt: bestAttempt, StudentName: device.StudentName, ClassName: device.ClassName})
		}
	}
	sort.SliceStable(boardAttempts, func(i, j int) bool {
		return boardAttempts[i].TopAttempt.Score > boardAttempts[j].TopAttempt.Score
	})
	return &models.LeaderBoard{Level: level, LeaderBoardStats: boardAttempts}
}


func GetClassLeaderBoard(service db.DatabaseService, level string, limit int, deviceId string) *models.LeaderBoard {
	return &models.LeaderBoard{}
}
