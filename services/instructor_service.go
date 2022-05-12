package services

import (
	"context"
	"efieldrestful/db"
	"efieldrestful/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const instructorCollection = "instructors"

func GetAllInstructors(service db.DatabaseService) []models.Instructor {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	mongoResult := service.FindAll(instructorCollection)
	defer mongoResult.Close(ctx)
	defer cancel()
	instructorList := make([]models.Instructor, 0)
	for mongoResult.Next(ctx) {
		instructor := models.Instructor{}
		err := mongoResult.Decode(&instructor)
		checkError(err)
		instructorList = append(instructorList, instructor)
	}
	return instructorList
}

func GetInstructorsByInstitution(service db.DatabaseService, institution string) []models.Instructor {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	mongoResult := service.FindByFieldMatchesString(instructorCollection, "institution", institution)
	defer mongoResult.Close(ctx)
	defer cancel()
	instructorList := make([]models.Instructor, 0)
	for mongoResult.Next(ctx) {
		instructor := models.Instructor{}
		err := mongoResult.Decode(&instructor)
		checkError(err)
		instructorList = append(instructorList, instructor)
	}
	return instructorList
}

func GetInstructorFromId(service db.DatabaseService, instructorId string) *models.Instructor {
	objectId, err := primitive.ObjectIDFromHex(instructorId)
	if err != nil {
		return nil
	}
	result := service.FieldMatchesString(instructorCollection, "_id", objectId)
	if result.Err() != nil {
		return nil
	} else {
		instructor := models.Instructor{}
		err := result.Decode(&instructor)
		checkError(err)
		return &instructor
	}
}

func StoreInstructor(service db.DatabaseService, instructor *models.Instructor) *models.Instructor {
	result := service.FieldMatchesString(instructorCollection, "email", instructor.Email)
	if result.Err() != nil {
		return nil
	} else {
		instructor := models.Instructor{}
		err := result.Decode(&instructor)
		checkError(err)
		return &instructor
	}
}

func DeleteInstructorById(service db.DatabaseService, instructorId string) *mongo.DeleteResult {
	return service.DeleteOneByFieldMatches(instructorCollection, "_id", instructorId)
}
