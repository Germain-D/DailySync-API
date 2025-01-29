package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetWeather(w http.ResponseWriter, r *http.Request, latitude, longitude string) {

	// URL de l'API Open-Meteo avec les paramètres
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&daily=weather_code,temperature_2m_max,temperature_2m_min,uv_index_max&timezone=auto&forecast_days=3", latitude, longitude)

	// Envoyer une requête GET
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la requête API : %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Vérifier le code de statut de la réponse
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("API renvoie un statut non-OK: %s", resp.Status), resp.StatusCode)
		return
	}

	// Lire le corps de la réponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture de la réponse : %v", err), http.StatusInternalServerError)
		return
	}

	// Décoder la réponse JSON dans une structure Go
	var weatherData map[string]interface{}
	if err := json.Unmarshal(body, &weatherData); err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du décodage JSON : %v", err), http.StatusInternalServerError)
		return
	}

	// Renvoyer les données météo au client en JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(weatherData); err != nil {
		log.Printf("Erreur lors de l'encodage JSON : %v", err)
	}
}
