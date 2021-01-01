package config

import (
	"log"
	"os"
	"os/user"

	"github.com/marcelovicentegc/kontrolio-cli/messages"
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
	filename := messages.KONTROLIO_CONFIG_FILENAME
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
	dbName := messages.KONTROLIO_DB_FILENAME
	return homePath + dbName
}
