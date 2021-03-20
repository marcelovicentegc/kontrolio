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
	filename := KontrolioConfigFilename
	return homePath + filename
}

func checkConfigFile() {
	filePath := getConfigFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		setNetworkMode(NetworkMode{Offline, ConfigIsMissing, NA})
		return
	}

	config := &Config{}

	config = GetConfig()

	if config.ApiKey == "" {
		setNetworkMode(NetworkMode{Offline, APIKeyIsMissing, NA})
	}
}

// GetLocalDataStorePath returns Kontrolio's
// local database path.
func GetLocalDataStorePath() string {
	homePath := getHomePath()
	dbName := KontrolioDatabaseFilename
	return homePath + dbName
}
