package models

type LeaderBoard struct {
	Level string `json:"level,omitempty" bson:"level,omitempty"`
	LeaderBoardStats []LeaderBoardStat `json:"stats,omitempty" bson:"stats,omitempty"`
}

type LeaderBoardStat struct {
	TopAttempt Attempt `json:"top_attempt,omitempty" bson:"top_attempt,omitempty"`
	StudentName string `json:"student_name,omitempty" bson:"student_name,omitempty"`
	ClassName string `json:"class_name,omitempty" bson:"class_name,omitempty"`
}

