package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// TideData représente une entrée de données de marée
type TideData struct {
	Height float64 `json:"height"`
	Time   string  `json:"time"`
	Type   string  `json:"type"`
}

// Station représente les informations de la station
type Station struct {
	Distance float64 `json:"distance"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	Name     string  `json:"name"`
	Source   string  `json:"source"`
}

// Meta représente les métadonnées de la réponse
type Meta struct {
	Cost         int     `json:"cost"`
	DailyQuota   int     `json:"dailyQuota"`
	Datum        string  `json:"datum"`
	End          string  `json:"end"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	Offset       int     `json:"offset"`
	RequestCount int     `json:"requestCount"`
	Start        string  `json:"start"`
	Station      Station `json:"station"`
}

// TideResponse représente la réponse complète de l'API
type TideResponse struct {
	Data []TideData `json:"data"`
	Meta Meta       `json:"meta"`
}

func GetTideState(w http.ResponseWriter, r *http.Request, StormApi string, lat string, lng string) {

	today := time.Now()
	json_name := "data/tide_" + today.Format("2006-01-02") + ".json"

	// Vérifier si le fichier existe
	if _, err := os.Stat(json_name); err == nil {
		// Le fichier existe
		file, err := os.Open(json_name)
		if err != nil {
			log.Fatalf("Erreur lors de l'ouverture du fichier : %v", err)
		}
		defer file.Close()

		w.Header().Set("Content-Type", "application/json")
		_, err = io.Copy(w, file)
		if err != nil {
			log.Fatalf("Erreur lors de la copie du fichier : %v", err)
		}

	} else {

		// Calculer les timestamps pour `start` et `end`
		now := time.Now()
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC) // Début de la journée
		end := start.AddDate(0, 0, 1)                                                // Début du jour suivant

		// Construction de l'URL avec les paramètres
		url := fmt.Sprintf("https://api.stormglass.io/v2/tide/extremes/point?lat=%s&lng=%s&start=%d&end=%d",
			lat, lng, start.Unix(), end.Unix())

		// Création de la requête HTTP
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("Erreur lors de la création de la requête : %v", err)
		}

		// Ajout des en-têtes
		req.Header.Add("Authorization", StormApi)

		// Envoi de la requête
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Erreur lors de l'envoi de la requête : %v", err)
		}
		defer resp.Body.Close()

		// Lecture de la réponse
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Erreur lors de la lecture de la réponse : %v", err)
		}

		fmt.Println("Réponse reçue :", string(body))

		// Décodage de la réponse JSON dans les structures Go
		var tideResponse TideResponse
		if err := json.Unmarshal(body, &tideResponse); err != nil {
			log.Fatalf("Erreur lors du décodage JSON : %v", err)
		}

		// créer un fichier json
		file, err := os.Create("data/tide_" + today.Format("2006-01-02") + ".json")
		if err != nil {
			log.Fatalf("Erreur lors de la création du fichier : %v", err)
		}
		defer file.Close()

		// Encoder les données JSON dans le fichier
		encoder := json.NewEncoder(file)
		if err := encoder.Encode(tideResponse); err != nil {
			log.Fatalf("Erreur lors de l'encodage JSON : %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = io.Copy(w, file)
		if err != nil {
			log.Fatalf("Erreur lors de la copie du fichier : %v", err)
		}

	}

}
