package common

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type configuration struct {
	Server, MongoDBHost, DBUser, DBpwd, Database string
}

type (
	appError struct {
		Error      string `json:"error"`
		Message    string `json:"message"`
		HttpStatus int    `json:"status"`
	}
	errorResoucre struct {
		Data appError `json:"data"`
	}
)

// AppConfig holds the configuration values from config.json file
var AppConfig configuration

// Initialize AppConfig
func initConfig() {

	loadAppConfig()
}

// Reads config.json and decode into AppConfig
func loadAppConfig() {
	file, err := os.Open("common/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig]: %s\n", err)
	}
}

//DisplayAppError used to return error message to client
func DisplayAppError(w http.ResponseWriter, handleError error, message string, code int) {
	errObj := appError{
		Error:      handleError.Error(),
		Message:    message,
		HttpStatus: code,
	}
	log.Printf("[AppError]: %s\n", handleError)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if j, err := json.Marshal(errorResoucre{Data: errObj}); err == nil {
		w.Write(j)
	}
}
