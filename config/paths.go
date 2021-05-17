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

// GetConfigFilePath returns the file path
// for .kontrolio.yaml configuration file
func GetConfigFilePath() string {
	homePath := getHomePath()
	filename := KontrolioConfigFilename
	return homePath + filename
}

// GetLocalDataStorePath returns Kontrolio's
// local database path.
func GetLocalDataStorePath() string {
	homePath := getHomePath()
	dbName := KontrolioDatabaseFilename
	return homePath + dbName
}
