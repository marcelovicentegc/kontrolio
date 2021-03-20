package cli

import (
	"fmt"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

func logs() {
	if config.Network.Status == config.Offline {
		fmt.Println(utils.YOURE_OFFLINE)

		records := db.GetOfflineRecords()

		var parsedRecords []utils.Record
		var currentDay *time.Time
		var log []string
		workWindowNanoseconds := int64(0)
		workNanoseconds := int64(0)

		for _, serializedRecord := range records {
			record := utils.DeserializeOfflineRecord(serializedRecord)
			parsedRecords = append(parsedRecords, record)
		}

		for index, parsedRecord := range parsedRecords {
			isLastRecord := index + 1 == len(parsedRecords)

			if currentDay == nil {
				endOfDay := utils.EndOfDay(parsedRecord.Time)
				currentDay = &endOfDay
				log = append(log, utils.FormatLogMessageHeader(currentDay))
			}

			if parsedRecord.Time.After(*currentDay) {
				log = append(log, 
					utils.FormatLogMessageFooter(time.Duration(workNanoseconds).String(), time.Duration(workWindowNanoseconds).String(),
				))
				
				// Resets time accumulators
				workWindowNanoseconds = 0
				workNanoseconds = 0

				endOfDay := utils.EndOfDay(parsedRecord.Time)
				currentDay = &endOfDay
				log = append(log, utils.FormatLogMessageHeader(currentDay))
			}

			if (!isLastRecord && parsedRecords[index + 1].Time.Before(*currentDay)) {
				workWindowNanoseconds = workWindowNanoseconds + utils.SubtractTime(parsedRecord.Time, parsedRecords[index + 1].Time)
			}
			
			log = append(log, fmt.Sprintln(parsedRecord.Time.Format(time.RFC3339) + " " + parsedRecord.Type))
			
			// We compute worked hours from records of type
			// "out"
			if (parsedRecord.Type == db.RecordTypeRegistry.Out) {
				workNanoseconds = workNanoseconds + utils.SubtractTime(parsedRecords[index - 1].Time, parsedRecord.Time)
			}

			// This condition means that we reached the last log,
			// thus we must append the accumulated time for the last 
			// day here
			if (isLastRecord) {
				// Covers the cases where the client has punched in but haven't
				// punched out yet, so we compute how much time has passed
				// between when it punched in and now.
				if (parsedRecord.Type == db.RecordTypeRegistry.In) {
					workNanoseconds= workNanoseconds + utils.SubtractTime(parsedRecord.Time, time.Now())
				}

				log = append(log, 
					utils.FormatLogMessageFooter(time.Duration(workNanoseconds).String(), time.Duration(workWindowNanoseconds).String(),
				))
			}
		}
		
		for _, l := range log {
			fmt.Print(l)
		}

		fmt.Println()
	}
}
