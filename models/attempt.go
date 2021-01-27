package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Attempt struct {
	AttemptId      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Level          string             `json:"level,omitempty" bson:"level,omitempty"`
	Track          string             `json:"track,omitempty" bson:"track,omitempty"`
	StarsCollected uint8              `json:"stars_collected,omitempty" bson:"stars_collected,omitempty"`
	Score          uint16             `json:"score,omitempty" bson:"score,omitempty"`
	Time           uint32             `json:"time,omitempty" bson:"time,omitempty"`
	Timestamp      string             `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
