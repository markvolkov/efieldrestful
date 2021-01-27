package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Device struct {
	DeviceId primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ClassId primitive.ObjectID `json:"class_id,omitempty" bson:"class_id,omitempty"`
	StudentName string `json:"student_name,omitempty" bson:"student_name,omitempty"`
	Attempts []Attempt `json:"attempts,omitempty" bson:"attempts,omitempty"`
}

