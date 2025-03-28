package config

import "os"

type Configuration struct {
	HostAddress string
	HostPort    string
}

func GetConfigurations() Configuration {
	return Configuration{
		HostAddress: getHostAddress(),
		HostPort:    getHostPort(),
	}
}

// get configuration from local environment variables or use the default

func getHostAddress() string {
	value := os.Getenv(hostAddress)
	if isEmptyOrWhiteSpace(value) {
		return "go-tiny.com"
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
