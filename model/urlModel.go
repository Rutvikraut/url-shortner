package model

type URL struct {
	ID           string `json:"id"`
	OriginalURL  string `json:"original_url"`
	ShortenedURL string `json:"shortened_url"`
}

var UrlDB = make(map[string]URL)

var DomainMetrics = make(map[string]int)