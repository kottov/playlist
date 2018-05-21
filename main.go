package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

var configPath = "config/config.json"

type Configuration struct {
	Port     string
	ClientId string
	UserId   string
}

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
	config := retrieveConfig(configPath)
	http.HandleFunc("/", foo)
	log.Printf("Web server started on port :%s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, nil))
}

func foo(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s, %s, %s", r.Method, r.Host, r.URL)

	profile := Profile{"Kostya", []string{"music", "codding"}}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

func retrieveConfig(filePath string) Configuration {
	file, _ := os.Open(filePath)
	defer file.Close()

	decoder := json.NewDecoder(file)

	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Println("error:", err)
	}
	log.Printf("Configuration retrieved from %s", filePath)
	return configuration
}
