package db

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

type recordTypeRegistry struct {
	In  string
	Out string
}

func newRecordTypeRegistry() *recordTypeRegistry {
	return &recordTypeRegistry{
		In:  "IN",
		Out: "OUT",
	}
}

var RecordTypeRegistry = newRecordTypeRegistry()

func openOrCreateDb() *os.File {
	filePath := config.GetLocalDataStorePath()

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)

		if err != nil {
			log.Fatalln("failed to create .kontrolio.csv", err)
		}

		writer := csv.NewWriter(file)
		defer writer.Flush()

		if err := writer.Write([]string{"Time", "Punched"}); err != nil {
			log.Fatalln("error writing header to .kontrolio.csv", err)
		}

		return file
	}

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalln("failed to open .kontrolio.csv", err)
	}

	return file
}

func readData() ([][]string, error) {
	file := openOrCreateDb()

	defer file.Close()

	reader := csv.NewReader(file)

	// Skip first line
	if _, err := reader.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := reader.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}

func GetRecords() []utils.Record {
	records, err := readData()

	if err != nil {
		log.Fatal(err)
	}

	var parsedRecords []utils.Record

	for _, record := range records {
		time, err := time.Parse(time.RFC3339, record[0])

		if err != nil {
			log.Fatal(err)
		}

		parsedRecord := utils.Record{
			Time: time,
			Type: record[1],
		}

		parsedRecords = append(parsedRecords, parsedRecord)
	}

	return parsedRecords
}

func SaveRecord() {
	var serializedRecord []string

	serializedRecord = append(serializedRecord, time.Now().In(time.Local).Format(time.RFC3339))

	currentData, err := readData()

	lastRecord := currentData[len(currentData)-1]

	if lastRecord[1] == RecordTypeRegistry.In {
		serializedRecord = append(serializedRecord, RecordTypeRegistry.Out)
	} else {
		serializedRecord = append(serializedRecord, RecordTypeRegistry.In)
	}

	if err != nil {
		log.Fatal(err)
	}

	headers := []string{"Time", "Punched"}
	currentData = append([][]string{headers}, currentData...)
	currentData = append(currentData, serializedRecord)

	file, err := os.Create(config.GetLocalDataStorePath())

	if err != nil {
		log.Fatalln("failed to create .kontrolio.csv", err)
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.WriteAll(currentData); err != nil {
		log.Fatalln("error writing header to .kontrolio.csv", err)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages.FormatPunchMessage(serializedRecord[1]))

}
