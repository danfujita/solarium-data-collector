package config_reader

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Token           string
	InfluxAddress   string
	InfluxUserName  string
	InfluxPassword  string
	InfluxDatabase  string
	InfluxTableName string
}

func Config() Configuration {
	file, _ := os.Open("config/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}
