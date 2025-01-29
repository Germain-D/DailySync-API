package main

import (
	"log"
	"net/http"

	"dailysync.com/handlers"
	"dailysync.com/middleware"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Appliquer le middleware d'authentification
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	// Définir les routes
	api.HandleFunc("/weather", handlers.GetWeather).Methods("GET")
	api.HandleFunc("/surf", handlers.GetSurfConditions).Methods("GET")
	api.HandleFunc("/tide", handlers.GetTideState).Methods("GET")
	api.HandleFunc("/party", handlers.GetTodaysParty).Methods("GET")
	api.HandleFunc("/btc", handlers.GetBTCPrice).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
