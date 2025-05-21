package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

var jokes = []string{
	"パンはパンでも食べられないパンは？ → フライパン！",
	"布団が吹っ飛んだ！",
	"トイレに行っといれ。",
	"犬がいるんだワン。",
	"象がぞうっとする話。",
}

var mu sync.Mutex

func main() {
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", NewRouter())
}

func NewRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/joke", getJokesHandler)
	return r
}

func getJokesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}
