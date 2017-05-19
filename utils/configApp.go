package utils

import (
	"os"
	"fmt"
	"encoding/json"
	"../models/config"
)

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	if err != nil { fmt.Println(err.Error()) }
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config) return config
}