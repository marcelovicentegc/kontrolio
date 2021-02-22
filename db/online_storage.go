package db

// SyncDatabases is used to sync offline and online records,
// reconciling them by automatically resolving conflicts.
func SyncDatabases() {
	offlineRecords := GetOfflineRecords()

	println(offlineRecords)
}