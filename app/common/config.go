package common

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	Email          string
	Database       string
	Role           string
	TwitterKey     string
	TwitterSecret  string
	FacebookKey    string
	FacebookSecret string
}

var AppConfig configuration

func InitConfig() {
	loadConfig()

}

func loadConfig() {
	file, err := os.Open("app/common/configurations.json")
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}

	json.NewDecoder(file).Decode(&AppConfig)
}
