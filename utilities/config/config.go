package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config stores application configuration file
type Config struct {
	DbConnectionString string `json:"connectionString,omitempty"`
	TokenSecret        string `json:"tokenSecret,omitempty"`
}

var config Config

// InitConfig opens the configuration file and load it into the Config object
func InitConfig() {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(configFile, &config); err != nil {
		panic("Unable to parse config file")
	}
}

// GetConfig returns the application configuration
func GetConfig() *Config {
	if !isConfigFilled() {
		InitConfig()
	}
	return &config
}

func isConfigFilled() bool {
	if (Config{}) == config {
		return false
	}
	return true
}
