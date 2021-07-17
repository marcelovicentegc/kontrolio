package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	KontrolioHealthCheckLocal          = "http://localhost:3000/api/ht"
	KontrolioHealthCheck               = "https://kontrolio.com/api/ht"
	KontrolioConfigFilename            = ".kontrolio.yaml"
	KontrolioDatabaseFilename          = ".kontrolio.csv"
	KontrolioGoogleTokenFilename       = ".kontrolio.google.token.json"
	KontrolioGoogleCredentialsFilename = ".kontrolio.google.credentials.json"
)

type Config struct {
	ApiKey      string   `yaml:"api_key"`
	Dev         string   `yaml:"dev"`
	Name        string   `yaml:"name"`
	WorkingDays []string `yaml:"working_days"`
	WorkingTime string   `yaml:"working_time"`
}

// GetConfig finds and reads the
// configuration file, returning
// its parsed data
func GetConfig() *Config {
	filePath := GetConfigFilePath()

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			fmt.Println(err)

			return nil
		}

		log.Fatal(err)
	}

	config := &Config{}

	err = yaml.Unmarshal(buf, config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}
