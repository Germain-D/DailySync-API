package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func GetTodaysParty(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	monthStr := fmt.Sprintf("%d", int(now.Month()))
	dayStr := fmt.Sprintf("%d", now.Day())

	file, err := os.Open("data/saints.json")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var data map[string]map[string][]string
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	saints, exists := data[monthStr][dayStr]
	if !exists || len(saints) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "No saints found for this day"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]string{"saints": saints})
}
