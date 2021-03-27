package config

const (
	// BaseURL is the base URL from Kontrolio's public API gateway.
	BaseURL = "https://kontrolio.com/api/"
	// LocalBaseURL is the base URL from Kontrolio's public API gateway
	// when it's running locally along this CLI client.
	LocalBaseURL = "http://localhost:3000/api/"
	// RecordEndpoint is the endpoint used for creating new records.
	RecordEndpoint = "record"
	// RecordsEndpoint is the endpoint used for querying records through
	// a filter.
	RecordsEndpoint = "records"
	// AllRecordsEndpoint is the endpoint used for fetching every record
	// ever saved on the remote database.
	AllRecordsEndpoint = "records/all"
)

// GetBaseURL is responsible for getting the
// development or production base URL.
func GetBaseURL() string {
	if IsDevEnvironment() {
		return LocalBaseURL
	} else {
		return BaseURL
	}
}
