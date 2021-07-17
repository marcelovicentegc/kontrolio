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

func getFile(filename string) string {
	homePath := getHomePath()
	return homePath + filename
}

// GetConfigFilePath returns the file path
// for .kontrolio.yaml configuration file
func GetConfigFilePath() string {
	return getFile(KontrolioConfigFilename)
}

// GetLocalDataStorePath returns Kontrolio's
// local database path.
func GetLocalDataStorePath() string {
	return getFile(KontrolioDatabaseFilename)
}

func GetGoogleTokenPath() string {
	return getFile(KontrolioGoogleTokenFilename)
}

func GetGoogleCredentialsPath() string {
	return getFile(KontrolioGoogleCredentialsFilename)
}
