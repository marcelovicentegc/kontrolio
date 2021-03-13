package config

import (
	"net/http"
)

// NetworkMode holds the current status
// of which the application is being ran,
// as well as the reason for failures on
// if it's in production.
type NetworkMode struct {
	Status int
	Reason int
	IsProduction int
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
	IS_PROD_ENVIRONMENT = iota
)

var NETWORK_MODE = NetworkMode{OFFLINE, NA, NA}

func setNetworkMode(networkMode NetworkMode) {
	NETWORK_MODE = networkMode
}

func checkConnection() {
	var healthCheckEndpoint string
	var isProd int

	if (IsDevEnvironment()) {
		isProd = NA
		healthCheckEndpoint = KONTROLIO_HEALTH_CHECK_LOCAL
	} else {
		isProd = IS_PROD_ENVIRONMENT
		healthCheckEndpoint = KONTROLIO_HEALTH_CHECK
	}

	res, err := http.Get(healthCheckEndpoint)

	if err != nil {
		setNetworkMode(NetworkMode{OFFLINE, NETWORK_IS_DOWN, isProd})
		return
	}

	if res.StatusCode == 200 {
		setNetworkMode(NetworkMode{ONLINE, NA, isProd})
		return
	}

	setNetworkMode(NetworkMode{OFFLINE, SERVICE_IS_DOWN, isProd})
}
