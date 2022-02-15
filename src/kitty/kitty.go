package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Cat struct {
	Name string `json:"name"`
}

func helloKittyHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello Kitty!")
	if err != nil {
		return
	}
}

func catAPIHandler(w http.ResponseWriter, r *http.Request) {
	cats := make([]Cat, 1)
	cats[0] = Cat{Name: "Ginger"}
	err := json.NewEncoder(w).Encode(cats)
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", helloKittyHandler)
	http.HandleFunc("/api/cats", catAPIHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
