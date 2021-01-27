package models

type LeaderBoard struct {
	Level string `json:"level,omitempty" bson:"level,omitempty"`
	Track string `json:"track,omitempty" bson:"track,omitempty"`
	LeaderBoardStats []LeaderBoardStat `json:"stats, omitempty" bson:"stats,omitempty"`
}

type LeaderBoardStat struct {
	TopAttempt Attempt `json:"top_attempt,omitempty" bson:"top_attempt,omitempty"`
	ClassName string `json:"class_name,omitempty" bson:"class_name,omitempty"`
}

