package config

// IsDevEnvironment returns true in case it's a dev environment,
// and false in case it's a production environment.
func IsDevEnvironment() bool {
	config := GetConfig()

	return config.Dev == "true"
}