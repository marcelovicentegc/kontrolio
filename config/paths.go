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
	filename := ".kontrolio.yaml"
	return homePath + filename
}

func checkConfigFileExistence() {
	filePath := getConfigFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		setNetworkMode(NetworkMode{OFFLINE, CONFIG_IS_MISSING})
		return
	}
}

func GetLocalDataStorePath() string {
	homePath := getHomePath()
	dbName := ".kontrolio.db"
	return homePath + dbName
}
