package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const (
	KONTROLIO_HEALTH_CHECK_LOCAL = "http://localhost:3000/api/ht"
	KONTROLIO_HEALTH_CHECK    = "https://kontrolio.com/api/ht"
	KONTROLIO_CONFIG_FILENAME = ".kontrolio.yaml"
	KONTROLIO_DB_FILENAME     = ".kontrolio.db"
)

type Config struct {
	ApiKey string `yaml:"api_key"`
	Dev string `yaml:"dev"`
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
		checkConfigFile()
	}
}
