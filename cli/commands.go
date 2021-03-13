package cli

import (
	"fmt"
	"log"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/client"
	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

func punch() {
	if config.NETWORK_MODE.Status == config.OFFLINE {
		fmt.Println(utils.YOURE_OFFLINE)
		recordType := db.SaveOfflineRecord()
		fmt.Println(utils.FormatPunchMessage(recordType))
		return
	}

	if config.NETWORK_MODE.Status == config.ONLINE {
		appConfig := config.GetConfig()
		utils.DisplayOnlineMessage(*appConfig)

		// TODO: Sync remote and local data, then save record remotely and locally.

		recordType := client.CreateRecord(appConfig.ApiKey)
		fmt.Println(utils.FormatPunchMessage(recordType))
	}
}

func workdayStatus(calledAlone bool) {
	today := utils.BeginningOfDay(time.Now())
	tomorrow := today.AddDate(0, 0, 1)

	if config.NETWORK_MODE.Status == config.OFFLINE {
		var todaysRecords []string
		var nanoseconds int64

		records := db.GetOfflineRecords()
		for _, serializedRecord := range records {
			record := utils.DeserializeOfflineRecord(serializedRecord)
			if record.Time.After(today) && record.Time.Before(tomorrow) {
				todaysRecords = append(todaysRecords, serializedRecord)
			}
		}

		for index, serializedTodaysRecord := range todaysRecords {
			record := utils.DeserializeOfflineRecord(serializedTodaysRecord)

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
				openedRecordIndex := utils.IndexOf(records, serializedTodaysRecord)
				openedRecord := utils.DeserializeOfflineRecord(records[openedRecordIndex])

				if openedRecord.Type == db.RecordTypeRegistry.In {
					nanoseconds = utils.SubtractTime(record.Time, openedRecord.Time)
				}

				continue
			}

			if record.Type == db.RecordTypeRegistry.Out {
				openedRecord := utils.DeserializeOfflineRecord(todaysRecords[index-1])
				nanoseconds = nanoseconds + utils.SubtractTime(openedRecord.Time, record.Time)

				continue
			}
		}

		fmt.Print(utils.WORKDAY_STATUS)
		fmt.Println(time.Duration(nanoseconds).String() + "\n")

		return
	}

	if config.NETWORK_MODE.Status == config.ONLINE {
		appConfig := config.GetConfig()
		if calledAlone {
			utils.DisplayOnlineMessage(*appConfig)
		}

		// TODO: Get workday status from remote database.
	}
}

func sync() {
	if config.NETWORK_MODE.Status == config.OFFLINE {
		if config.NETWORK_MODE.Reason == config.NETWORK_IS_DOWN {
			log.Fatal(utils.SYNC_OFFLINE)
		}

		if config.NETWORK_MODE.Reason == config.SERVICE_IS_DOWN {
			log.Fatal(utils.SYNC_SERVICE_DOWN)
		}

		if config.NETWORK_MODE.Reason == config.CONFIG_IS_MISSING {
			log.Fatal(utils.SYNC_CONFIG_MISSING)
		}

		if config.NETWORK_MODE.Reason == config.API_KEY_IS_MISSING {
			log.Fatal(utils.SYNC_CONFIG_MISSING)
		}
	}

	appConfig := config.GetConfig()

	utils.DisplayOnlineMessage(*appConfig)

	offlineRecords := db.GetOfflineRecords()

	var parsedOfflineRecords []utils.Record

	for _, serializedOfflineRecord := range offlineRecords {
		record := utils.DeserializeOfflineRecord(serializedOfflineRecord)
		parsedOfflineRecords = append(parsedOfflineRecords, record)
	}

	parsedOfflineRecords = utils.ReverseRecords(parsedOfflineRecords)

	println(parsedOfflineRecords)

	onlineRecords := client.GetAllRecords(appConfig.ApiKey)

	println(onlineRecords)

	// TODO: Sync records
}


func getLogs() {
	if config.NETWORK_MODE.Status == config.OFFLINE {
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
