package config

import (
	"encoding/json"
	"fmt"
	"mailer/model"
	"os"
)

// Init configs
func Init() model.Config {
	file, _ := os.Open("./config.json")
	// Close file after usage
	defer file.Close()
	var c model.Config
	if err := json.NewDecoder(file).Decode(&c); err != nil {
		fmt.Println("Cannot read config file:", err)
		os.Exit(1)
	}
	return c
}
