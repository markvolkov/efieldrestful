package main

/**
    ===================
	Author: Mark Volkov
    ===================
 */

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	TIMEOUT = 10
)

/**
    =======================
	Type Struct Definitions
    =======================
 */

type AppConfig struct {
	AppPort    string `json:"port"`
	AppAddress string `json:"address"`
	MongoURI   string `json:"uri"`
	DbName     string `json:"dbName"`
}

func readValues(env string) AppConfig {
	envFile, err := os.Open("./config/config." + strings.ToLower(env) + ".json")
	checkError(err)
	appConfig := AppConfig{}
	json.NewDecoder(envFile).Decode(&appConfig)
	return appConfig
}

type App struct {
	Config      AppConfig
	Router      *mux.Router
	MongoClient *mongo.Client
}

func (app *App) init(env string) {
	app.Config = readValues(env)
	app.Router = mux.NewRouter().StrictSlash(true)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(app.Config.MongoURI))
	checkError(err)
	app.MongoClient = client
	app.setUpRoutes()
}

func (app *App) runApplication() {
	log.Println("Listening @ http://" + app.Config.AppAddress + ":" + app.Config.AppPort)
	server := &http.Server{
		Handler:      app.Router,
		Addr:         app.Config.AppAddress + ":" + app.Config.AppPort,
		WriteTimeout: (TIMEOUT * 1.5) * time.Second,
		ReadTimeout:  (TIMEOUT * 1.5) * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

type Attempt struct {
	AttemptId      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Level          string             `json:"level,omitempty" bson:"level,omitempty"`
	StarsCollected uint8              `json:"stars_collected,omitempty" bson:"stars_collected,omitempty"`
	Score          uint16             `json:"score,omitempty" bson:"score,omitempty"`
	Time           uint32             `json:"time,omitempty" bson:"time,omitempty"`
	Timestamp      string             `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}

type Device struct {
	DeviceId string `json:"_id,omitempty" bson:"_id,omitempty"`
	//StudentName string `json:"studentName,omitempty" bson:"studentName,omitempty"` TODO: Not sure if we need this.
	Attempts []Attempt `json:"attempts,omitempty" bson:"attempts,omitempty"`
}

type Class struct {
	ClassId string   `json:"_id,omitempty" bson:"_id,omitempty"`
	Devices []Device `json:"devices,omitempty" bson:"devices,omitempty"`
}

/**
    ========================
	Json Encoders / Decoders
    ========================
 */

func decodeAttempt(r *http.Request) *Attempt {
	var toDecode Attempt
	json.NewDecoder(r.Body).Decode(&toDecode)
	return &toDecode
}

func decodeClass(r *http.Request) *Class {
	var toDecode Class
	json.NewDecoder(r.Body).Decode(&toDecode)
	return &toDecode
}

func encodeError(w http.ResponseWriter, error string) {
	err := map[string]string{"error": error}
	json.NewEncoder(w).Encode(&err)
}

/**
    ========================
	MongoDB Helper Functions
    ========================
 */

func (app *App) getCollection(collection string) *mongo.Collection {
	return app.MongoClient.Database(app.Config.DbName).Collection(collection)
}

func (app *App) insertOne(collection string, payload interface{}) *mongo.InsertOneResult {
	log.Println("Inserting payload into collection " + collection)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	result, err := app.getCollection(collection).InsertOne(ctx, payload)
	checkError(err)
	return result
}

func (app *App) findAll(collection string) *mongo.Cursor {
	log.Println("Finding all from collection: " + collection)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	filter := bson.M{}
	findOpts := options.FindOptions{}
	findOpts.SetMaxTime(time.Second * (TIMEOUT / 5))
	result, err := app.getCollection(collection).Find(ctx, filter, &findOpts)
	checkError(err)
	return result
}

func (app *App) getById(collection string, bytes []byte) *mongo.SingleResult {
	log.Println("Getting by id from " + collection + " with _id " + string(bytes))
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	filter := bson.M{"_id": string(bytes)}
	findOpts := options.FindOneOptions{}
	findOpts.SetMaxTime(time.Second * (TIMEOUT / 5))
	result := app.getCollection(collection).FindOne(ctx, filter, &findOpts)
	return result
}

/***
    ======================
	Http Handler Functions
    ======================

	createClass - Will create a class with a specified "class-id"
	getClass - Will retrieve a class by a "class-id"
	classList - Will return all the current classes from the database
	storeAttempt - Will store an attempt from a student to a specified class
    getAttemptsByDeviceId - Will receive attempts from a device based on its "device-id"
    getDevices - Will return all the current devices from the database
 */

func (app *App) createClass(w http.ResponseWriter, r *http.Request) {
	classPayload := decodeClass(r)
	mongoResult := app.getById("classes", []byte(classPayload.ClassId))
	if mongoResult.Err() != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		result := app.insertOne("classes", classPayload)
		json.NewEncoder(w).Encode(&result)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		encodeError(w, "A Class With That ID Already Exists.")
	}
}

func (app *App) getClass(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	mongoResult := app.getById("classes", []byte(params["classId"]))
	if mongoResult.Err() != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		encodeError(w, "A Class With That ID Doesn't Exists.")
	} else {
		class := Class{}
		err := mongoResult.Decode(&class)
		checkError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(class)
	}
}

func (app *App) classList(w http.ResponseWriter, r *http.Request) {
	mongoResult := app.findAll("classes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer mongoResult.Close(ctx)
	classList := make([]Class, 0)
	for mongoResult.Next(ctx) {
		currClass := Class{}
		err := mongoResult.Decode(&currClass)
		checkError(err)
		classList = append(classList, currClass)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(classList)
}

func (app *App) storeAttempt(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deviceId := params["deviceId"]
	mongoResult := app.getById("devices", []byte(deviceId))
	device := Device{DeviceId: deviceId, Attempts: make([]Attempt, 0)}
	if mongoResult.Err() == nil {
		mongoResult.Decode(&device)
	}
	attempt := decodeAttempt(r)
	device.Attempts = append(device.Attempts, *attempt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	result := app.insertOne("devices", device)
	json.NewEncoder(w).Encode(&result)
}

func (app *App) getAttemptsByDeviceId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	deviceId := params["deviceId"]
	mongoResult := app.getById("devices", []byte(deviceId))
	if mongoResult.Err() != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		encodeError(w, "A Device With That ID Doesn't Exists.")
	} else {
		device := Device{}
		err := mongoResult.Decode(&device)
		checkError(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(device)
	}
}

func (app *App) getDevices(w http.ResponseWriter, r *http.Request) {
	mongoResult := app.findAll("devices")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer mongoResult.Close(ctx)
	devices := make([]Device, 0)
	for mongoResult.Next(ctx) {
		currDevice := Device{}
		mongoResult.Decode(&currDevice)
		devices = append(devices, currDevice)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(devices)
}

func (app *App) setUpRoutes() {
	app.Router.HandleFunc("/class/", app.classList).Methods("GET")
	app.Router.HandleFunc("/class/", app.createClass).Methods("POST")
	app.Router.HandleFunc("/class/{classId}", app.getClass).Methods("GET")
	app.Router.HandleFunc("/device/", app.getDevices).Methods("GET")
	app.Router.HandleFunc("/device/{deviceId}/", app.storeAttempt).Methods("POST")
	app.Router.HandleFunc("/device/{deviceId}/", app.getAttemptsByDeviceId).Methods("GET")
}

func main() {
	envFlag := flag.String("env", "dev", "Your environment config to run: ( dev || prod )")
	app := App{}
	app.init(*envFlag)
	app.runApplication()
	defer func() {
		ctx, _ := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
		if err := app.MongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Fatal error: %s", err.Error())
	}
}
