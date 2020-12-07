package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Class struct {
	ClassId primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	ClassName string `json:"className,omitempty" bson:"className,omitempty"`
	AccessCode string `json:"accessCode,omitempty" bson:"accessCode,omitempty"`
	Devices []Device `json:"devices,omitempty" bson:"devices,omitempty"`
}