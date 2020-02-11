package main

import (
	"encoding/json"
	"log"
	"os"
)

type provider struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	RedirectURL  string `json:"RedirectUrl"`
}

// Configuration struct represents top level app-config
// see config.json.example
type Configuration struct {
	Secret string              `json:"secretKey"`
	Oauth  map[string]provider `json:"oauth"`
}

// NewConfig function returns app configuration
func NewConfig() (Configuration, error) {
	file, err := os.Open("./config.json")
	if err != nil {
		log.Printf("Error reading config.json %s\n", err)
		return Configuration{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Println("error: ", err)
		return Configuration{}, err
	}

	return config, nil
}
