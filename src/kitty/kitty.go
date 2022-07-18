package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Cat struct {
	Name string `json:"name"`
}

func helloKittyHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello Kitty!\n")
	if err != nil {
		return
	}
}

func catAPIHandler(w http.ResponseWriter, r *http.Request) {
	cats := make([]Cat, 2)
	cats[0] = Cat{Name: "Ginger"}
	cats[1] = Cat{Name: "Sheila"}
	err := json.NewEncoder(w).Encode(cats)
	if err != nil {
		return
	}
}

func main() {
	const portNumber = 8080
	fmt.Println("Starting Cats and Kittens server on localhost:" + strconv.Itoa(portNumber) + " ...")
	http.HandleFunc("/", helloKittyHandler)
	http.HandleFunc("/api/cats", catAPIHandler)
	err := http.ListenAndServe(":"+strconv.Itoa(portNumber), nil)
	if err != nil {
		return
	}
}
