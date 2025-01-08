package handlers

import (
	"net/http"
	"url-shortner/utils"
)

func RedirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := utils.GetUrl(id)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusNotFound)
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
