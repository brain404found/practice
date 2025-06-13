package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"rate-limiter/limiter"
	"time"

	"github.com/brianvoe/gofakeit"
)

var (
	port     = ":8080"
	limit    = 10
	interval = time.Second
)

func GetNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	}{
		Id:    gofakeit.Number(1, 1000),
		Title: gofakeit.Sentence(10),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	ctx := context.Background()
	limiter := limiter.NewRateLimiter(ctx, limit, interval)

	http.HandleFunc("/note", func(w http.ResponseWriter, r *http.Request) {
		// As middleware pre-process the request
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		GetNote(w, r)
	})

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO:
// - add tests for the rate limiter
// - add metrics for requests
// - prometheus config
// - docker-compose file
// - request generator for load testing and POC rate limit
// - describe the pattern in detail in README.md
