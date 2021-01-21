package config

import (
	"net/http"
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
	SERVICE_IS_DOWN    = iota
	NETWORK_IS_DOWN    = iota
	CONFIG_IS_MISSING  = iota
	API_KEY_IS_MISSING = iota
	NA                 = iota
)

var NETWORK_MODE = NetworkMode{OFFLINE, NA}

func setNetworkMode(networkMode NetworkMode) {
	NETWORK_MODE = networkMode
}

func checkConnection() {
	res, err := http.Get(KONTROLIO_HEALTH_CHECK)
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
