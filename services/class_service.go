package services

import (
	"context"
	"efieldrestful/app"
	"efieldrestful/db"
	"efieldrestful/models"
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
	 result := service.FieldMatchesString(classesCollection, "accessCode", accessCode)
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
	 result := service.FieldMatchesString(classesCollection, "_id", classId)
	 if result.Err() != nil {
	 	return nil
	 } else {
	 	class := models.Class{}
	 	err := result.Decode(&class)
	 	checkError(err)
	 	return &class
	 }
}

func GetDevicesFromClass(service db.DatabaseService, classId string) []models.Device {
	return GetClassFromId(service, classId).Devices
}

func StoreDeviceInClass(service db.DatabaseService, classId string, deviceId string) {
	class := GetClassFromId(service, classId)
	deviceToAdd := GetDeviceById(service, deviceId)
	class.Devices = append(class.Devices, *deviceToAdd)
	service.InsertOne(classesCollection, class)
}

func DeleteClassByAccessCode(app app.App, accessCode string) {
	app.DatabaseService.DeleteOneByFieldMatches(classesCollection, "accessCode", accessCode)
}

func DeleteClassById(app app.App, classId string) {
	app.DatabaseService.DeleteOneByFieldMatches(classesCollection, "_id", classId)
}
