package utls

import (
	"OnlineShop/models"
	"encoding/json"
	"log"
	"os"
)

var Config models.Configuration

func LoadConfiguration() (string, string) {
	initKeys()
	file, err := os.Open("utls/config.json")

	if err != nil {
		log.Println("Cannot open config file:", err)
	}
	decoder := json.NewDecoder(file)

	Config = models.Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Println("Failed to decode the file", err)
	}

	file.Close()

	return Config.ServerIp, Config.ServerPort
}

