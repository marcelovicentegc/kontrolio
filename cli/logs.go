package cli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

func logs(tail *string) {
	records := db.GetRecords()

	var currentDay *time.Time
	var log []string
	workWindowNanoseconds := int64(0)
	workNanoseconds := int64(0)

	for index, parsedRecord := range records {
		isLastRecord := index+1 == len(records)

		if currentDay == nil {
			endOfDay := utils.EndOfDay(parsedRecord.Time)
			currentDay = &endOfDay
			log = append(log, messages.FormatLogMessageHeader(currentDay))
		}

		if parsedRecord.Time.After(*currentDay) {
			log = append(log,
				messages.FormatLogMessageFooter(time.Duration(workNanoseconds).String(), time.Duration(workWindowNanoseconds).String()))

			// Resets time accumulators
			workWindowNanoseconds = 0
			workNanoseconds = 0

			endOfDay := utils.EndOfDay(parsedRecord.Time)
			currentDay = &endOfDay
			log = append(log, messages.FormatLogMessageHeader(currentDay))
		}

		if !isLastRecord && records[index+1].Time.Before(*currentDay) {
			workWindowNanoseconds = workWindowNanoseconds + utils.SubtractTime(parsedRecord.Time, records[index+1].Time)
		}

		log = append(log, messages.FormatLogMessage(parsedRecord))

		// We compute worked hours from records of type
		// "out"
		if parsedRecord.Type == db.RecordTypeRegistry.Out {
			workNanoseconds = workNanoseconds + utils.SubtractTime(records[index-1].Time, parsedRecord.Time)
		}

		// This condition means that we reached the last log,
		// thus we must append the accumulated time for the last
		// day here
		if isLastRecord {
			// Covers the cases where the client has punched in but haven't
			// punched out yet, so we compute how much time has passed
			// between when it punched in and now.
			if parsedRecord.Type == db.RecordTypeRegistry.In {
				workNanoseconds = workNanoseconds + utils.SubtractTime(parsedRecord.Time, time.Now())
			}

			log = append(log,
				messages.FormatLogMessageFooter(time.Duration(workNanoseconds).String(), time.Duration(workWindowNanoseconds).String()))
		}
	}

	if tail != nil {
		if tail, err := strconv.Atoi(*tail); err == nil && tail > 0 {
			var indices []int

			for index, l := range log {
				if strings.Contains(l, "\r") {
					indices = append(indices, index)
				}
			}

			if tail > len(indices) {
				for _, l := range log {
					fmt.Print(l)
				}

				fmt.Println()

				return
			}

			end := len(log)
			start := indices[len(indices)-tail]

			for _, l := range log[start:end] {
				fmt.Print(l)
			}

			fmt.Println()

			return
		}

		fmt.Println("Invalid `tail` value, printing all logs")
	}

	for _, l := range log {
		fmt.Print(l)
	}

	fmt.Println()
}
