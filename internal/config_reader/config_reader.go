package config_reader

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Configuration struct {
	TokenSecret     string
	InfluxAddress   string
	InfluxUserName  string
	InfluxPassword  string
	InfluxDatabase  string
	InfluxTableName string
	S3BucketName    string
	PortNumber      int
}

var tokenSecret = os.Getenv("JWT_TOKEN")
var influxAddress = os.Getenv("INFLUX_ADDR")
var influxUserName = os.Getenv("INFLUX_USERNAME")
var influxPassword = os.Getenv("INFLUX_PW")
var influxDatabase = os.Getenv("INFLUX_DB")
var influxTableName = os.Getenv("INFLUX_TN")
var s3BucketName = os.Getenv("S3_BUCKET_NAME")
var portNumber = os.Getenv("PORT_NUMBER")
var configFile = os.Getenv("CONFIG_FILE")

func configFromfile(fileName string) Configuration {
	file, _ := os.Open("configs/" + fileName)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

func configFromArgs() Configuration {
	var configuration Configuration
	configuration.TokenSecret = tokenSecret
	configuration.InfluxAddress = influxAddress
	configuration.InfluxUserName = influxUserName
	configuration.InfluxPassword = influxPassword
	configuration.InfluxDatabase = influxDatabase
	configuration.InfluxTableName = influxTableName
	configuration.S3BucketName = s3BucketName
	strPortNumber, err := strconv.Atoi(portNumber)
	if err != nil {
		strPortNumber = 8080
	} else {
		configuration.PortNumber = strPortNumber
	}
	return configuration
}

func Config() Configuration {
	if configFile != "" {
		return configFromfile(configFile)
	} else {
		return configFromArgs()
	}
}
