package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
	log.Println("Server started!")
}

func foo(w http.ResponseWriter, r *http.Request) {
	log.Println("Request recieved!")
	profile := Profile{"Kostya", []string{"music", "codding"}}
	log.Printf("Profile created: %s", profile)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
	log.Printf("Response sent.")
}
