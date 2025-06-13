package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/brianvoe/gofakeit"
)

func main() {
	http.HandleFunc("/note", GetNote)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

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
