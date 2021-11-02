package config

import (
	"log"
	"os/user"
	"path/filepath"
)

func getHomePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir + "/"
}

func getFilePath(filename string) string {
	homePath := getHomePath()
	return homePath + filename
}

// GetConfigFilePath returns the file path
// for .kontrolio.yaml configuration file
func GetConfigFilePath() string {
	return getFilePath(KontrolioConfigFilename)
}

// GetLocalDataStorePath returns Kontrolio's
// local database path.
func GetLocalDataStorePath() string {
	return getFilePath(KontrolioDatabaseFilename)
}

// GetGoogleCredentialsPath returns Kontrolio's
// credentials.json file path.
func GetGoogleCredentialsPath() string {
	filePath, err := filepath.Abs("./" + KontrolioGoogleCredentialsFilename)
	if err != nil {
		log.Fatal(err)
	}

	return filePath
}
