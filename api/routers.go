package api

import (
	"github.com/gorilla/mux"
)

func BindRoutes(router *mux.Router) {
	router.HandleFunc("/ping", Ping).Methods("GET")
	router.HandleFunc("/deck/create", Create).Methods("GET")
	router.HandleFunc("/deck/open", Open).Methods("GET")
}
