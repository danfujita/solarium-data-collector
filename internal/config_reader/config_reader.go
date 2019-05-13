package config_reader

import (
	"encoding/json"
	"flag"
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
	S3BucketName    string
	PortNumber      int
}

var token = flag.String("token", "", "JWT token secret")
var influxAddress = flag.String("influx-addr", "", "Address of influx db")
var influxUserName = flag.String("influx-uname", "", "User name of influx db")
var influxPassword = flag.String("influx-pw", "", "Password of influx db")
var influxDatabase = flag.String("influx-db", "", "Database of influx db")
var influxTableName = flag.String("influx-tn", "", "Table name of influx db")
var s3BucketName = flag.String("s3-bucket", "", "S3 bucket Name")
var portNumber = flag.Int("port-num", 8080, "Port Number")
var configFile = flag.String("config-file", "", "File name")

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
	configuration.Token = *token
	configuration.InfluxAddress = *influxAddress
	configuration.InfluxUserName = *influxUserName
	configuration.InfluxPassword = *influxPassword
	configuration.InfluxDatabase = *influxDatabase
	configuration.InfluxTableName = *influxTableName
	configuration.S3BucketName = *s3BucketName
	configuration.PortNumber = *portNumber
	return configuration
}

func Config() Configuration {
	flag.Parse()
	if *configFile != "" {
		fmt.Printf(*configFile)
		return configFromfile(*configFile)
	} else {
		return configFromArgs()
	}
}
