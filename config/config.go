package config

import (
	"encoding/json"
	"fmt"
	"mailer/model"
	"os"
)

// Config of the eapplication
var Config model.Config

// Init configs
func Init() model.Config {
	// Read config
	file, _ := os.Open("./config.json")
	defer file.Close()
	err := json.NewDecoder(file).Decode(&Config)
	if err != nil {
		fmt.Println("Cannot read config file:", err)
		os.Exit(1)
	}
	return Config
}
