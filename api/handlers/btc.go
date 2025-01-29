package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Struct pour parser la réponse JSON
type BTCResponse struct {
	Time struct {
		Updated string `json:"updated"`
	} `json:"time"`
	Bpi struct {
		USD struct {
			Rate string `json:"rate"`
		} `json:"USD"`
		EUR struct {
			Rate string `json:"rate"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func GetBTCPrice(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Getting BTC price\n")

	// Faire la requête à l'API CoinDesk
	resp, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Vérifier le statut de la réponse
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch BTC price", resp.StatusCode)
		return
	}

	// Lire le corps de la réponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parser la réponse JSON dans la structure BTCResponse
	var btcResp BTCResponse
	if err := json.Unmarshal(body, &btcResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Créer un JSON simplifié avec les champs souhaités
	simplifiedResponse := map[string]string{
		"updated":     btcResp.Time.Updated,
		"rate_in_usd": btcResp.Bpi.USD.Rate,
		"rate_in_eur": btcResp.Bpi.EUR.Rate,
	}

	// Convertir le JSON simplifié en bytes
	jsonResponse, err := json.Marshal(simplifiedResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Envoyer la réponse
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
