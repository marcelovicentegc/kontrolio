package utils

import (
	"log"
	"strings"
	"time"
)

type Record struct {
	Time time.Time
	Type string
}

func SerializeOfflineRecord(record Record) ([]byte, []byte) {
	return []byte(record.Time.String()), []byte(record.Type)
}

func DeserializeOfflineRecord(record string) Record {
	splitRecord := strings.Split(record, ",")
	time, err := time.Parse(`"`+time.RFC3339+`"`, splitRecord[0])

	if err != nil {
		log.Fatal(err)
	}

	return Record{Time: time, Type: splitRecord[1]}
}
