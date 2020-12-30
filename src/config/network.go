package config

import (
	"net/http"
	"os"
)

type NetworkMode struct {
	Status int
	Reason int
}

const (
	OFFLINE = iota
	ONLINE  = iota
)

const (
	SERVICE_IS_DOWN   = iota
	NETWORK_IS_DOWN   = iota
	CONFIG_IS_MISSING = iota
	NA                = iota
)

var NETWORK_MODE = NetworkMode{OFFLINE, NA}

func setNetworkMode(networkMode NetworkMode) {
	NETWORK_MODE = networkMode
}

func checkConnection() {
	res, err := http.Get("https://api.kontrolio.com")
	if err != nil {
		setNetworkMode(NetworkMode{OFFLINE, NETWORK_IS_DOWN})
		return
	}

	if res.StatusCode == 200 {
		setNetworkMode(NetworkMode{ONLINE, NA})
		return
	}

	setNetworkMode(NetworkMode{OFFLINE, SERVICE_IS_DOWN})
}

func checkConfigFileExistence() {
	filePath := getConfigFilePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		setNetworkMode(NetworkMode{OFFLINE, CONFIG_IS_MISSING})
		return
	}
}

func ConfigNetworkMode() {
	checkConnection()

	if NETWORK_MODE.Status == ONLINE {
		checkConfigFileExistence()
	}
}
