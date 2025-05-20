package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
)

var jokes = []string{
	"I'm reading a book on anti-gravity. It’s impossible to put down!",
	"Did you hear about the restaurant on the moon? Great food, no atmosphere.",
	"Why don’t skeletons fight each other? They don’t have the guts.",
}

var mu sync.Mutex

func main() {
	http.ListenAndServe(":8080", NewRouter())
}

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/joke", getJokeHandler)
	r.Post("/joke", postJokeHandler)
	return r
}

func getJokeHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	rand.Seed(time.Now().UnixNano())
	joke := jokes[rand.Intn(len(jokes))]
	json.NewEncoder(w).Encode(map[string]string{"joke": joke})
}

func postJokeHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Joke string `json:"joke"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Joke == "" {
		http.Error(w, "Invalid joke", http.StatusBadRequest)
		return
	}
	mu.Lock()
	jokes = append(jokes, body.Joke)
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "received"})
}
