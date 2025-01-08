package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"url-shortner/model"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	type DomainMetric struct {
		Domain string
		Count  int
	}

	var metrics []DomainMetric
	for domain, count := range model.DomainMetrics {
		metrics = append(metrics, DomainMetric{Domain: domain, Count: count})
	}
	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].Count > metrics[j].Count
	})

	if len(metrics) > 3 {
		metrics = metrics[:3]
	}

	response := struct {
		Metrics []DomainMetric `json:"metrics"`
	}{Metrics: metrics}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
