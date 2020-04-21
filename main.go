package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/deck-of-cards/src/api"
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