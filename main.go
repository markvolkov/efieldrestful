package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

const newLine = "\n"

/**
	Type Struct Definitions
 */

type Attempt struct {
	DeviceId string `json:"deviceId"`
	StudentName string `json:"studentName"`
	Level string `json:"level"`
	StarsCollected uint8 `json:"starsCollected"`
	Score uint16 `json:"score"`
	Time uint32 `json:"time"`
}

type Class struct {
	ClassId string `json:"classId"`
	Attempts []Attempt `json:"attempts"`
}

func encodeAttempt(w http.ResponseWriter, toEncode Attempt) {
	json.NewEncoder(w).Encode(toEncode)
}

func encodeClass(w http.ResponseWriter, toEncode Class) {
	json.NewEncoder(w).Encode(toEncode)
}

func decodeAttempt(r *http.Request) Attempt {
	var toDecode Attempt
	json.NewDecoder(r.Body).Decode(&toDecode)
	return toDecode
}

func decodeClass(r *http.Request) Class {
	var toDecode Class
	json.NewDecoder(r.Body).Decode(&toDecode)
	return toDecode
}

/***
	Http Handler Functions

	createClass - Will create a class with a specified "class-id"
	getClass - Will retrieve a class by a "class-id"
	storeAttempt - Will store an attempt from a student to a specified class
    getAttemptsByDeviceId - Will receive attempts from a device based on its "device-id"

 */

func createClass(w http.ResponseWriter, r *http.Request) {

}

func getClass(w http.ResponseWriter, r *http.Request) {

}

func storeAttempt(w http.ResponseWriter, r *http.Request) {

}

func getAttemptsByDeviceId(w http.ResponseWriter, r *http.Request) {
}


func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	mux := http.NewServeMux()
	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}

