package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ApiKey string `yaml:"api_key"`
}

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

	if NETWORK_MODE.Status == ONLINE {
		checkConfigFileExistence()
	}
}
