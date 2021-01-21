package config

import (
	"log"
	"os"
	"os/user"
)

func getHomePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir + "/"
}

func getConfigFilePath() string {
	homePath := getHomePath()
	filename := KONTROLIO_CONFIG_FILENAME
	return homePath + filename
}

func checkConfigFile() {
	filePath := getConfigFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		setNetworkMode(NetworkMode{OFFLINE, CONFIG_IS_MISSING})
		return
	}

	config := &Config{}

	config = GetConfig()

	if config.ApiKey == "" {
		setNetworkMode(NetworkMode{OFFLINE, API_KEY_IS_MISSING})
	}
}

func GetLocalDataStorePath() string {
	homePath := getHomePath()
	dbName := KONTROLIO_DB_FILENAME
	return homePath + dbName
}
