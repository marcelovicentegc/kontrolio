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
	Offline = iota
	Online = iota
)

const (
	ServiceIsDown    = iota
	NetworkIsDown    = iota
	ConfigIsMissing  = iota
	APIKeyIsMissing = iota
	NA                 = iota
	IsProdEnvironment = iota
)

var Network = NetworkMode{Offline, NA, NA}

func setNetworkMode(networkMode NetworkMode) {
	Network = networkMode
}

func checkConnection() {
	var healthCheckEndpoint string
	var isProd int

	if (IsDevEnvironment()) {
		isProd = NA
		healthCheckEndpoint = KontrolioHealthCheckLocal
	} else {
		isProd = IsProdEnvironment
		healthCheckEndpoint = KontrolioHealthCheck
	}

	res, err := http.Get(healthCheckEndpoint)

	if err != nil {
		setNetworkMode(NetworkMode{Offline, NetworkIsDown, isProd})
		return
	}

	if res.StatusCode == 200 {
		setNetworkMode(NetworkMode{Online, NA, isProd})
		return
	}

	setNetworkMode(NetworkMode{Offline, ServiceIsDown, isProd})
}
