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
	fmt.Fprintf(w, "Hello Kitty!")
}

func catAPIHandler(w http.ResponseWriter, r *http.Request) {
	cats := make([]Cat, 1)
	cats[0] = Cat{Name: "Ginger"}
	json.NewEncoder(w).Encode(cats)
}

func main() {
	http.HandleFunc("/", helloKittyHandler)
	http.HandleFunc("/api/cats", catAPIHandler)
	http.ListenAndServe(":8080", nil)
}
