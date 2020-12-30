package config

import (
	"io/ioutil"
	"log"
	"os/user"

	"gopkg.in/yaml.v2"
)

func getHomePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir + "/"
}

type config struct {
	ApiKey string `yaml:"api_key"`
}

func GetConfig() *config {
	homePath := getHomePath()
	filename := ".kontrolio.yaml"
	filePath := homePath + filename

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	config := &config{}
	err = yaml.Unmarshal(buf, config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}
