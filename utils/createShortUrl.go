package utils

import (
	"url-shortner/model"
)

func CreateShortUrl(originalURL string) string {

	for _, url := range model.UrlDB {
		if url.OriginalURL == originalURL {
			return url.ShortenedURL
		}
	}

	shortUrl := generateShortUrl(originalURL)
	id := shortUrl
	model.UrlDB[id] = model.URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortenedURL: shortUrl,
	}

	domain := ExtractDomain(originalURL)
	model.DomainMetrics[domain]++

	return shortUrl
}
