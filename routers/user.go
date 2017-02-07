package routers

import (
	"github.com/golang-es/go-cms/controllers"

	"github.com/gorilla/mux"
)

func SetUserRouter(router *mux.Router) *mux.Router {
	//Rutas públicas para registro y login de usuarios
	router.HandleFunc("/register", controllers.ShowRegister).Methods("GET")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	return router
}
