package config

import (
	"log"
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

func GetLocalDataStorePath() string {
	homePath := getHomePath()
	dbName := ".kontrolio.db"
	return homePath + dbName
}
