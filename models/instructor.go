package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Instructor struct {
	InstructorId primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Institution string `json:"institution,omitempty" bson:"institution,omitempty"`
	Classes []Class `json:"classes,omitempty'" bson:"classes,omitempty"`
}
