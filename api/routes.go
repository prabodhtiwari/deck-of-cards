package api

import (
	"github.com/gorilla/mux"
	"github.com/deck-of-cards/api/handlers"
)

func BindRoutes(router *mux.Router) {
	router.HandleFunc("/ping", handlers.Ping).Methods("GET")
	router.HandleFunc("/deck/create", handlers.Create).Methods("GET")
	router.HandleFunc("/deck/open", handlers.Open).Methods("GET")
	router.HandleFunc("/deck/draw", handlers.Draw).Methods("GET")
}
