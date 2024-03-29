package cli

import (
	"fmt"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

func status(calledAlone bool) {
	today := utils.BeginningOfDay(time.Now())
	tomorrow := today.AddDate(0, 0, 1)

	var todaysRecords []utils.Record
	var nanoseconds int64

	records := db.GetRecords()

	for _, record := range records {
		if record.Time.After(today) && record.Time.Before(tomorrow) {
			todaysRecords = append(todaysRecords, record)
		}
	}

	for index, record := range todaysRecords {
		// Covers the cases where the client has punched in but haven't
		// punched out yet, so we compute how much time has passed
		// between when it punched in and now.
		if index == (len(todaysRecords)-1) && record.Type == db.RecordTypeRegistry.In {
			nanoseconds = nanoseconds + utils.SubtractTime(record.Time, time.Now())
			continue
		}

		// Covers the case where the client has punched in yesterday,
		// but punched out today. Example: you staretd working yesterday
		// @ 11PM, but only stopped working today @ 2AM.
		if index == 0 && record.Type == db.RecordTypeRegistry.Out {
			openedRecordIndex := utils.IndexOf(records, record)
			openedRecord := records[openedRecordIndex]

			if openedRecord.Type == db.RecordTypeRegistry.In {
				nanoseconds = utils.SubtractTime(record.Time, openedRecord.Time)
			}

			continue
		}

		if record.Type == db.RecordTypeRegistry.Out {
			openedRecord := todaysRecords[index-1]
			nanoseconds = nanoseconds + utils.SubtractTime(openedRecord.Time, record.Time)

			continue
		}
	}

	fmt.Println(messages.FormatStatusMessage(nanoseconds))
}
