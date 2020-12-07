package services

import (
	"context"
	"efieldrestful/db"
	"efieldrestful/models"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

const (
	TIMEOUT = 10
	devicesCollection = "devices"
)

func GetAllDevices(service db.DatabaseService) []models.Device {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	mongoResult := service.FindAll(devicesCollection)
	defer mongoResult.Close(ctx)
	defer cancel()
	deviceList := make([]models.Device, 0)
	for mongoResult.Next(ctx) {
		device := models.Device{}
		mongoResult.Decode(&device)
		deviceList = append(deviceList, device)
	}
	return deviceList
}

func GetDevicesByStudentName(service db.DatabaseService, studentName string) []models.Device {
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	mongoResult := service.FindByFieldMatchesString(devicesCollection, "student_name", studentName)
	defer mongoResult.Close(ctx)
	defer cancel()
	deviceList := make([]models.Device, 0)
	for mongoResult.Next(ctx) {
		device := models.Device{}
		mongoResult.Decode(&device)
		deviceList = append(deviceList, device)
	}
	return deviceList
}

func GetDeviceById(service db.DatabaseService, id string) *models.Device {
	result := service.FieldMatchesString(devicesCollection, "_id", id)
	if result.Err() != nil {
		return nil
	}
	device := models.Device{}
	err := result.Decode(&device)
	checkError(err)
	return &device
}

func StoreAttemptFromDevice(service db.DatabaseService, deviceId string, attempt models.Attempt) *mongo.InsertOneResult {
	device := GetDeviceById(service, deviceId)
	if device != nil {
		device.Attempts = append(device.Attempts, attempt)
		return service.InsertOne(devicesCollection, device)
	} else {
		log.Println("There device was not found, could not store attempt for device id: " + deviceId)
		return nil
	}
}

func DeleteDeviceById(service db.DatabaseService, deviceId string) {
	service.DeleteOneByFieldMatches(deviceId, "_id", deviceId)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
