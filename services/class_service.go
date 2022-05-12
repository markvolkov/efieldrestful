package services

import (
	"context"
	"efieldrestful/db"
	"efieldrestful/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const classesCollection = "classes"

func GetAllClasses(service db.DatabaseService) []models.Class {
	 ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	 mongoResult := service.FindAll(classesCollection)
	 defer mongoResult.Close(ctx)
	 defer cancel()
	 classList := make([]models.Class, 0)
	 for mongoResult.Next(ctx) {
	 	class := models.Class{}
	 	err := mongoResult.Decode(&class)
	 	checkError(err)
	 	classList = append(classList, class)
	 }
	 return classList
}

func SaveClass(service db.DatabaseService, classPayload *models.Class) *mongo.InsertOneResult {
	mongoResult := service.InsertOne(classesCollection, classPayload)
	return mongoResult
}

func GetClassFromAccessCode(service db.DatabaseService, accessCode string) *models.Class {
	 result := service.FieldMatchesString(classesCollection, "access_code", accessCode)
	 if result.Err() != nil {
	 	return nil
	 } else {
	 	class := models.Class{}
	 	err := result.Decode(&class)
	 	checkError(err)
	 	return &class
	 }
}

func GetClassFromId(service db.DatabaseService, classId string) *models.Class {
	 objectId, err := primitive.ObjectIDFromHex(classId)
	 if err != nil {
	 	return nil
	 }
	 result := service.FieldMatchesString(classesCollection, "_id", objectId)
	 if result.Err() != nil {
	 	return nil
	 } else {
	 	class := models.Class{}
	 	err := result.Decode(&class)
	 	checkError(err)
	 	return &class
	 }
}

func GetDevicesFromClass(service db.DatabaseService, classId string) []string {
	return GetClassFromId(service, classId).Devices
}

func StoreDeviceInClass(service db.DatabaseService, classId string, deviceId string) *mongo.UpdateResult {
	class := GetClassFromId(service, classId)
	class.Devices = append(class.Devices, deviceId)
	return service.UpdateOne(classesCollection, bson.M{"_id": classId },  class)
}

func DeleteClassByAccessCode(service db.DatabaseService, accessCode string) *mongo.DeleteResult {
	return service.DeleteOneByFieldMatches(classesCollection, "access_code", accessCode)
}

func DeleteClassById(service db.DatabaseService, classId string) *mongo.DeleteResult {
	objectId, err := primitive.ObjectIDFromHex(classId)
	if err != nil {
		return nil
	}
	return service.DeleteOneByFieldMatches(classesCollection, "_id", objectId)
}
