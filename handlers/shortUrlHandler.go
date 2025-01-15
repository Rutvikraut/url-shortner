package handlers

import (
	"encoding/json"
	"net/http"
	"url-shortner/utils"
)

func ShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)

		return
	}

	shortUrl_ := utils.CreateShortUrl(data.URL)

	response := struct {
		ShortURL string `json:"short_url`
	}{ShortURL: shortUrl_}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}