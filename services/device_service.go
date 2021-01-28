package services

import (
	"context"
	"efieldrestful/db"
	"efieldrestful/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	result := service.FieldMatchesString(devicesCollection, "_id", objectId)
	if result.Err() != nil {
		return nil
	}
	device := models.Device{}
	err = result.Decode(&device)
	checkError(err)
	return &device
}


func StoreDevice(service db.DatabaseService, device *models.Device) *mongo.InsertOneResult {
	log.Println("Storing new device")
	if GetDeviceById(service, device.DeviceId.Hex()) == nil {
		return service.InsertOne(devicesCollection, device)
	}
	return nil
}

func StoreAttemptFromDevice(service db.DatabaseService, deviceId string, attempt models.Attempt) *mongo.UpdateResult {
	device := GetDeviceById(service, deviceId)
	if device != nil {
		device.Attempts = append(device.Attempts, attempt)
		objectId, err := primitive.ObjectIDFromHex(deviceId)
		if err != nil {
			return nil
		}
		return service.UpdateOne(devicesCollection, bson.M{"_id": objectId }, device)
	} else {
		log.Println("There device was not found, could not store attempt for device id: " + deviceId)
		return nil
	}
}

func DeleteDeviceById(service db.DatabaseService, deviceId string) (error, *mongo.DeleteResult) {
	objectId, err := primitive.ObjectIDFromHex(deviceId)
	if err != nil {
		return err, nil
	}
	return nil, service.DeleteOneByFieldMatches(deviceId, "_id", objectId)
}

func checkError(err error) {
	if err != nil {
		log.Printf("Fatal error: %s\n", err.Error())
	}
}
