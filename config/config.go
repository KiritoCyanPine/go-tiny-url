package config

import (
	"fmt"
	"os"
)

var configInstance *Configuration

type Configuration struct {
	HostAddress string
	HostPort    string
}

func GetConfigurations() *Configuration {
	if configInstance == nil {
		configInstance = &Configuration{
			HostAddress: getHostAddress(),
			HostPort:    getHostPort(),
		}
	}

	fmt.Println("Loading configuration from environment...")
	fmt.Printf("\n%v\n", configInstance)

	return configInstance
}

// get configuration from local environment variables or use the default

func getHostAddress() string {
	value := os.Getenv(hostAddress)
	if isEmptyOrWhiteSpace(value) {
		return "http://localhost:8080"
	}

	return value
}

func getHostPort() string {
	value := os.Getenv(hostPort)
	if isEmptyOrWhiteSpace(value) {
		return ""
	}

	return value
}
