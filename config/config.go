package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const (
	KontrolioHealthCheckLocal = "http://localhost:3000/api/ht"
	KontrolioHealthCheck    = "https://kontrolio.com/api/ht"
	KontrolioConfigFilename = ".kontrolio.yaml"
	KontrolioDatabaseFilename     = ".kontrolio.db"
)

type Config struct {
	ApiKey string `yaml:"api_key"`
	Dev string `yaml:"dev"`
}

// GetConfig finds and reads the configuration file, returning
// its structured data
func GetConfig() *Config {
	filePath := getConfigFilePath()

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	config := &Config{}
	err = yaml.Unmarshal(buf, config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}

func ConfigNetworkMode() {
	checkConnection()

	if Network.Status == Online {
		checkConfigFile()
	}
}
