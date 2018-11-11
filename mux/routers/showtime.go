package routers

import (
	"github.com/gorilla/mux"
	"../controllers"
	"../authorization"
)

func SetShowTimeRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/showtimes", controllers.GetShowTimes).Methods("GET")
	router.Handle("/showtimes/{id}", authorization.AuthMiddleware(controllers.GetShowTimeById)).Methods("GET")
	return router
}
