package main

import (
	"log"
	"net/http"

	"github.com/ezechidc/taskmanager/keys"

	"github.com/codegangsta/negroni"
	"github.com/ezechidc/taskmanager/common"
	"github.com/ezechidc/taskmanager/routers"
)

func main() {
	//call generatekeys to generate key pair
	keys.GenerateKeys(2048)
	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create a negroni instance
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
