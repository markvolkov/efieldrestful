package app

import (
	"context"
	"efieldrestful/db"
	"efieldrestful/encrypt"
	"efieldrestful/resources"
	"encoding/json"
	"github.com/gorilla/mux"
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
	App Definitions
    =======================
 */

type AppConfig struct {
	AppPort    string `json:"port"`
	AppAddress string `json:"address"`
	MongoURI   string `json:"uri"`
	DbName     string `json:"dbName"`
	BasicUsername string `json:"basicUsername"`
	BasicPassword string `json:"basicPassword"`
}

func readValues(env string) AppConfig {
	envFile, err := os.Open("config/config." + strings.ToLower(env) + ".json")
	checkError(err)
	appConfig := AppConfig{}
	json.NewDecoder(envFile).Decode(&appConfig)
	appConfig.BasicUsername = string(encrypt.EncryptData([]byte(appConfig.BasicUsername)))
	appConfig.BasicPassword = string(encrypt.EncryptData([]byte(appConfig.BasicPassword)))
	return appConfig
}

type App struct {
	Config      AppConfig
	Router      *mux.Router
	db.DatabaseService
}

func (app *App) Init(env string) {
	app.Config = readValues(env)
	app.Router = mux.NewRouter().StrictSlash(true)
	ctx, cancel := context.WithTimeout(context.Background(), TIMEOUT*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(app.Config.MongoURI))
	checkError(err)
	app.DatabaseService = db.DatabaseService{Client: client, DbName: app.Config.DbName}
	app.setUpRoutes()
}

func (app *App) setUpRoutes() {
	//leaderboard handlers usage GET /leaderboard?limit=10&level=1&track=1&global=false/
	app.Router.HandleFunc("/leaderboard/", resources.GetLeaderBoard(app.DatabaseService)).Methods("GET", "OPTIONS")

	//class handlers
	app.Router.HandleFunc("/class/", resources.CreateClass(app.DatabaseService)).Methods("POST", "OPTIONS")
	app.Router.HandleFunc("/class/", resources.ClassList(app.DatabaseService)).Methods("GET", "OPTIONS")
	app.Router.HandleFunc("/class/{classId}", resources.GetClass(app.DatabaseService)).Methods("GET", "OPTIONS")

	//device handlers
	app.Router.HandleFunc("/device/", resources.StoreAttempt(app.DatabaseService)).Methods("POST", "OPTIONS")
	app.Router.HandleFunc("/device/", resources.GetDevices(app.DatabaseService)).Methods("GET", "OPTIONS")
	app.Router.HandleFunc("/device/{deviceId}/", resources.GetAttemptsByDeviceId(app.DatabaseService)).Methods("GET", "OPTIONS")

	//TODO: instructor handlers
	//app.Router.HandleFunc("/instructor/", app.getDevices).Methods("GET", "OPTIONS")
	//app.Router.HandleFunc("/instructor/{institution}/", app.storeAttempt).Methods("GET", "OPTIONS")
	//app.Router.HandleFunc("/instructor/{instructorId}/", app.getAttemptsByDeviceId).Methods("GET", "OPTIONS")

	//authentication middleware
	app.Router.Use(app.checkAuthenicationMiddleware)
}

func (app *App) checkAuthenicationMiddleware(nextRequest http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, UPDATE, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, Host, Accept, Authorization")

		if r.Method == "OPTIONS" {
			return
		}

		user, pass, _ := r.BasicAuth()
		authenticated := encrypt.CheckData([]byte(app.Config.BasicUsername), []byte(user)) && encrypt.CheckData([]byte(app.Config.BasicPassword), []byte(pass))
		if !authenticated {
			http.Error(w, "Forbidden: Must authenticate.", http.StatusForbidden)
		} else {
			nextRequest.ServeHTTP(w, r)
		}
	})
}

func (app *App) RunApplication() {
	log.Println("Listening @ http://" + app.Config.AppAddress + ":" + app.Config.AppPort)
	server := &http.Server{
		Handler:      app.Router,
		Addr:         app.Config.AppAddress + ":" + app.Config.AppPort,
		WriteTimeout: (TIMEOUT * 1.5) * time.Second,
		ReadTimeout:  (TIMEOUT * 1.5) * time.Second,
	}
	app.Router.Use(mux.CORSMethodMiddleware(app.Router))
	log.Fatal(server.ListenAndServe())
}

func checkError(err error) {
	if err != nil {
		log.Printf("Error: %s\n", err.Error())
	}
}

