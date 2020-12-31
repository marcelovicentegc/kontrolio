package utils

import (
	"log"
	"strings"
	"time"
)

func SerializeOfflineRecord(record string) (time.Time, string) {
	splitRecord := strings.Split(record, ",")
	time, err := time.Parse(`"`+time.RFC3339+`"`, splitRecord[0])

	if err != nil {
		log.Fatal(err)
	}

	return time, splitRecord[1]
}

func BeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}
