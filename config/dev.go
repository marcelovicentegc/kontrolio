package config

// IsDevEnvironment returns true in case it's a dev environment,
// and false in case it's a production environment based on the
// value set on the .kontrolio.yaml file.
func IsDevEnvironment() bool {
	config := GetConfig()

	return config.Dev == "true"
}
