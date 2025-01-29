package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func GetPageContent(url string) (string, error) {
	// Faire une requête HTTP pour obtenir le contenu de la page
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Lire le corps de la réponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convertir le corps en chaîne de caractères
	htmlContent := string(body)

	return htmlContent, nil
}

// Structures pour mapper le JSON
type Forecast struct {
	Key     int     `json:"key"`
	DateKey string  `json:"dateKey"`
	Values  []Value `json:"values"`
}

type Value struct {
	DateHour        string `json:"dateHour"`
	HoulePrimaire   string `json:"houlePrimaire"`
	HouleSecondaire string `json:"houleSecondaire"`
	HouleDirection  string `json:"houleDirection"`
	HoulePeriode    string `json:"houlePeriode"`
	Wind            string `json:"wind"`
	WindDirection   string `json:"windDirection"`
	Stars           string `json:"stars"`
}

func GetSurfConditions(w http.ResponseWriter, r *http.Request, url string) {

	// Récupérer le contenu de la page
	htmlContent, err := GetPageContent(url)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du contenu de la page:", err)
		return
	}

	// Regex pour extraire le contenu entre les balises <script> et </script>
	regex := regexp.MustCompile(`(?s)<script>(.*?)<\/script>`)
	matches := regex.FindAllStringSubmatch(string(htmlContent), -1)

	// Afficher les résultats
	if len(matches) == 0 {
		fmt.Println("Aucun contenu trouvé entre les balises <script> et </script>.")
		return
	}

	const regex2 = `var\s+data\s+=\s+(.*?);`

	var re = regexp.MustCompile(regex2)

	var data = `var data = {"foo": "bar"};`

	for _, match := range matches {

		result := re.FindStringSubmatch(match[1])
		if len(result) > 0 {
			data = result[1]

		}
	}
	// Using strings.TrimPrefix to remove "var data = " if it exists
	data = strings.TrimPrefix(data, "var data = ")

	// Remove any remaining whitespace
	data = strings.TrimSpace(data)

	var forecasts []Forecast
	err = json.Unmarshal([]byte(data), &forecasts)
	if err != nil {
		fmt.Println("Erreur lors du parsing JSON:", err)
		return
	}

	for i, forecast := range forecasts {

		for j, value := range forecast.Values {
			stars_count := 3
			stars_count -= strings.Count(value.Stars, "fa fa-star-o")
			forecasts[i].Values[j].Stars = fmt.Sprintf("%d", stars_count)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forecasts)

}
