package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Class struct {
	ClassId primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	ClassName string `json:"class_name,omitempty" bson:"class_name,omitempty"`
	AccessCode string `json:"access_code,omitempty" bson:"access_code,omitempty"`
	//TODO: Make this a db ref based on the _id
	Devices []string `json:"devices,omitempty" bson:"devices,omitempty"`
}