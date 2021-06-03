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

	var log []string

	logBuilder(records, func(workNanoseconds *int64, workWindowNanoseconds *int64, currentDay *time.Time, record *utils.Record) {
		if workNanoseconds != nil && workWindowNanoseconds != nil {
			log = append(log,
				messages.FormatLogMessageFooter(time.Duration(*workNanoseconds).String(), time.Duration(*workWindowNanoseconds).String()))
		} else if record == nil && currentDay != nil {
			log = append(log, messages.FormatLogMessageHeader(currentDay))
		} else {
			log = append(log, messages.FormatLogMessage(*record))
		}
	})

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

func logBuilder(records []utils.Record, logger func(workNanoseconds *int64, workWindowNanoseconds *int64, currentDay *time.Time, record *utils.Record)) {
	workWindowNanoseconds := int64(0)
	workNanoseconds := int64(0)
	var currentDay *time.Time

	for index, record := range records {
		isLastRecord := index+1 == len(records)

		if currentDay == nil {
			endOfDay := utils.EndOfDay(record.Time)
			currentDay = &endOfDay
			logger(nil, nil, currentDay, nil)
		}

		if record.Time.After(*currentDay) {
			logger(&workNanoseconds, &workWindowNanoseconds, nil, nil)

			// Resets time accumulators
			workWindowNanoseconds = 0
			workNanoseconds = 0

			logger(nil, nil, currentDay, nil)
		}

		if !isLastRecord && records[index+1].Time.Before(*currentDay) {
			workWindowNanoseconds = workWindowNanoseconds + utils.SubtractTime(record.Time, records[index+1].Time)
		}

		logger(nil, nil, nil, &record)

		// We compute worked hours from records of type
		// "out"
		if record.Type == db.RecordTypeRegistry.Out {
			workNanoseconds = workNanoseconds + utils.SubtractTime(records[index-1].Time, record.Time)
		}

		// This condition means that we reached the last log,
		// thus we must append the accumulated time for the last
		// day here
		if isLastRecord {
			// Covers the cases where the client has punched in but haven't
			// punched out yet, so we compute how much time has passed
			// between when it punched in and now.
			if record.Type == db.RecordTypeRegistry.In {
				workNanoseconds = workNanoseconds + utils.SubtractTime(record.Time, time.Now())
			}

			logger(&workNanoseconds, &workWindowNanoseconds, nil, nil)
		}
	}
}
