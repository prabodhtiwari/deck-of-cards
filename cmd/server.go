package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/deck-of-cards/api"
)

func main() {

	router := mux.NewRouter()
	router = initRoutes(router)
	addr := ":3000"
	log.Println("Server running at ", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Println(err)
	}

}

func initRoutes(router *mux.Router) *mux.Router {
	api.BindRoutes(router)
	return router
}