package db

import (
	"encoding/csv"
	"fmt"
	"io"
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

func readData() ([][]string, error) {
	dataStorePath := config.GetLocalDataStorePath()

	// Opens existing data store file or creates one
	file, err := os.OpenFile(dataStorePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("failed to open or create "+dataStorePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Write headers if file is empty
	if _, err := reader.Read(); err == io.EOF {
		writer := csv.NewWriter(file)
		defer writer.Flush()

		headers := []string{"Time", "Punched"}

		if err := writer.Write(headers); err != nil {
			log.Fatalln("error writing headers to "+dataStorePath, err)
		}

		return [][]string{}, nil
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

// SaveRecord saves a record on the punch command.
func SaveRecord() {
	var serializedRecord []string
	dataStorePath := config.GetLocalDataStorePath()

	serializedRecord = append(serializedRecord, time.Now().In(time.Local).Format(time.RFC3339))

	currentData, err := readData()

	if err != nil {
		log.Fatal(err)
	}

	currentDataLength := len(currentData)

	var lastRecord []string

	if currentDataLength > 0 {
		lastRecord = currentData[len(currentData)-1]

		if lastRecord[1] == RecordTypeRegistry.In {
			serializedRecord = append(serializedRecord, RecordTypeRegistry.Out)
		} else {
			serializedRecord = append(serializedRecord, RecordTypeRegistry.In)
		}
	} else {
		serializedRecord = append(serializedRecord, RecordTypeRegistry.In)
	}

	file, err := os.OpenFile(dataStorePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("failed to open or create "+dataStorePath, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write most recent installation status
	if err := writer.Write(serializedRecord); err != nil {
		log.Fatalln("error writing data to "+dataStorePath, err)
	}

	fmt.Println(messages.FormatPunchMessage(serializedRecord[1]))
}
