package main

import (
	"log"
	"net/http"

	"dailysync.com/handlers"
	"dailysync.com/middleware"
	"dailysync.com/utils"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Load environment variables
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize logger with the log level from environment variables
	err = utils.Initialize(config.LogLevel)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer utils.Sync()

	sugar := utils.SugaredLogger
	sugar.Infow("Starting application with config",
		"config", config,
	)

	// Appliquer le middleware d'authentification
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	// Définir les routes
	api.HandleFunc("/weather", handlers.GetWeather).Methods("GET")
	api.HandleFunc("/surf", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetSurfConditions(w, r, config.SurfReportLink)
	}).Methods("GET")
	api.HandleFunc("/tide", handlers.GetTideState).Methods("GET")
	api.HandleFunc("/party", handlers.GetTodaysParty).Methods("GET")
	api.HandleFunc("/btc", handlers.GetBTCPrice).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
