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

func SerializeOfflineRecord(record Record) ([]byte, []byte) {
	fmt.Println(record.Time.Format(time.RFC3339))
	return []byte(record.Time.Format(time.RFC3339)), []byte(record.Type)
}

func DeserializeOfflineRecord(record string) Record {
	splitRecord := strings.Split(record, ",")
	replacer := strings.NewReplacer("[", "", "]", "")
	time, err := time.Parse(time.RFC3339, replacer.Replace(splitRecord[0]))

	if err != nil {
		log.Fatal(err)
	}

	return Record{Time: time, Type: splitRecord[1]}
}
