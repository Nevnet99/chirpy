package main

import (
	"fmt"
	"net/http"
	"os"
)

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("./templates/metrics.html")

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(string(b), cfg.fileserverHits)))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}
