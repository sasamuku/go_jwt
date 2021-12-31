package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type post struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	URL    string `json:"url"`
}

func main() {
	r := mux.NewRouter()
	r.Handle("/public", public)
	r.Handle("/auth", GetTokenHandler)
	r.Handle("/private", JwtMiddleware.Handler(private))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

var public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title:  "Big News",
		Author: "Ken",
		URL:    "https://example.com/test",
	}
	json.NewEncoder(w).Encode(post)
})

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	post := &post{
		Title:  "Small News",
		Author: "John",
		URL:    "https://example.com/test",
	}
	json.NewEncoder(w).Encode(post)
})
