package configs

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Token string
	InfluxAddress string
	Username string
	Password string
}

func Config() Configuration{
	file, _ := os.Open("configs/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

