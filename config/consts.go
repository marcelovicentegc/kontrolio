package config

const (
	KontrolioHealthCheckLocal          = "http://localhost:3000/api"
	KontrolioHealthCheck               = "https://kontrolio.com/api"
	KontrolioConfigFilename            = ".kontrolio.yaml"
	KontrolioDatabaseFilename          = ".kontrolio.csv"
	KontrolioGoogleCredentialsFilename = "credentials.json"
	KontrolioGoogleScope               = "https://www.googleapis.com/auth/drive.file"
	GoogleDriveOAuthCallback           = "/oauth/drive/redirect"
	KontrolioFilename                  = "Kontrolio punch record"
	KontrolioFileMimeType              = "text/csv"
)
