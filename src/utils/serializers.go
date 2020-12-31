package utils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Record struct {
	Time time.Time
	Type string
}

func ByteSerializeOfflineRecord(record Record) ([]byte, []byte) {
	return []byte(record.Time.Format(time.RFC3339)), []byte(record.Type)
}

func SerializeOfflineRecord(key []byte, value []byte) string {
	return fmt.Sprintf("%s,%s", key, value)
}

func DeserializeOfflineRecord(record string) Record {
	splitRecord := strings.Split(record, ",")
	time, err := time.Parse(time.RFC3339, splitRecord[0])

	if err != nil {
		log.Fatal(err)
	}

	return Record{Time: time, Type: splitRecord[1]}
}
