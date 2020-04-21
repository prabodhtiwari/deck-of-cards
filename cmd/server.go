package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/deck-of-cards/api"
)

func main() {
	router := mux.NewRouter()
	api.BindRoutes(router)
	addr := ":3000"
	log.Println("Server running at...", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Println(err)
	}
}
