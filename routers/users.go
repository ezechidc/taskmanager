package routers

import (
	"github.com/ezechidc/taskmanager/controllers"
	"github.com/gorilla/mux"
)

//SetUserRoutes sets user routes for login and register
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
