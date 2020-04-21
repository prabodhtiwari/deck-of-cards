package api

import (
	"github.com/gorilla/mux"
)

func BindRoutes(router *mux.Router) {
	router.HandleFunc("/ping", Ping).Methods("GET")
}
