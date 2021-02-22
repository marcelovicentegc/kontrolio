package cli

import (
	"fmt"
	"log"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/client"
	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

func punch() {
	if config.NETWORK_MODE.Status == config.OFFLINE {
		fmt.Println(messages.YOURE_OFFLINE)
		recordType := db.SaveOfflineRecord()
		fmt.Println("Punched " + recordType + messages.PUNCH_SUCCESS + "\n")
		return
	}

	if config.NETWORK_MODE.Status == config.ONLINE {
		appConfig := config.GetConfig()
		utils.DisplayOnlineMessage(*appConfig)

		// TODO: Sync remote and local data, then save record remotely and locally.

		recordType := client.CreateRecord(appConfig.ApiKey)
		fmt.Println("Punched " + recordType + messages.PUNCH_SUCCESS + "\n")
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

		fmt.Print(messages.WORKDAY_STATUS)
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
			log.Fatal(messages.SYNC_OFFLINE)
		}

		if config.NETWORK_MODE.Reason == config.SERVICE_IS_DOWN {
			log.Fatal(messages.SYNC_SERVICE_DOWN)
		}

		if config.NETWORK_MODE.Reason == config.CONFIG_IS_MISSING {
			log.Fatal(messages.SYNC_CONFIG_MISSING)
		}

		if config.NETWORK_MODE.Reason == config.API_KEY_IS_MISSING {
			log.Fatal(messages.SYNC_CONFIG_MISSING)
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
		fmt.Println(messages.YOURE_OFFLINE)

		records := db.GetOfflineRecords()

		var parsedRecords []utils.Record
		var currentDay *time.Time

		for _, serializedRecord := range records {
			record := utils.DeserializeOfflineRecord(serializedRecord)
			parsedRecords = append(parsedRecords, record)
		}

		parsedRecords = utils.ReverseRecords(parsedRecords)

		for _, parsedRecord := range parsedRecords {
			if currentDay == nil {
				beginningOfDay := utils.BeginningOfDay(parsedRecord.Time)
				currentDay = &beginningOfDay
				fmt.Println(string(utils.ColorCyan), "\n"+currentDay.Format(time.RFC850), string(utils.ColorReset))
			}

			if parsedRecord.Time.Before(*currentDay) {
				beginningOfDay := utils.BeginningOfDay(parsedRecord.Time)
				currentDay = &beginningOfDay
				fmt.Println(string(utils.ColorCyan), "\n"+currentDay.Format(time.RFC850), string(utils.ColorReset))
			}

			fmt.Println(parsedRecord.Time.Format(time.RFC3339) + " " + parsedRecord.Type)
		}

		fmt.Println()
	}
}
