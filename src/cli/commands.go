package cli

import (
	"fmt"
	"log"
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/src/config"
	"github.com/marcelovicentegc/kontrolio-cli/src/db"
	"github.com/marcelovicentegc/kontrolio-cli/src/utils"
)

func punch() {
	if config.NETWORK_MODE.Status == config.OFFLINE {
		recordType := db.SaveOfflineRecord()
		fmt.Println(recordType + " sucessfully.")
		return
	}

	if config.NETWORK_MODE.Status == config.ONLINE {
		appConfig := config.GetConfig()
		fmt.Println("You're online with the " + appConfig.ApiKey + " API key!")

		// TODO: Sync remote and local data, then save record remotely and locally.
	}
}

func workdayStatus() {
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
			if index == (len(todaysRecords)-1) && record.Type == db.PUNCHED_IN {
				nanoseconds = nanoseconds + utils.SubtractTime(record.Time, time.Now())
				continue
			}

			// Covers the case where the client has punched in yesterday,
			// but punched out today. Example: you staretd working yesterday
			// @ 11PM, but only stopped working today @ 2AM.
			if index == 0 && record.Type == db.PUNCHED_OUT {
				openedRecordIndex := utils.IndexOf(records, serializedTodaysRecord)
				openedRecord := utils.DeserializeOfflineRecord(records[openedRecordIndex])

				if openedRecord.Type == db.PUNCHED_IN {
					nanoseconds = utils.SubtractTime(record.Time, openedRecord.Time)
				}

				continue
			}

			if record.Type == db.PUNCHED_OUT {
				openedRecord := utils.DeserializeOfflineRecord(todaysRecords[index-1])
				nanoseconds = nanoseconds + utils.SubtractTime(openedRecord.Time, record.Time)

				continue
			}
		}

		fmt.Print("\nToday, you've worked: ")
		fmt.Println(time.Duration(nanoseconds).String() + "\n")

		return
	}

	if config.NETWORK_MODE.Status == config.ONLINE {
		appConfig := config.GetConfig()
		fmt.Println("You're online with the " + appConfig.ApiKey + " API key!")

		// TODO: Get workday status from remote database.
	}
}

func sync() {
	if config.NETWORK_MODE.Status == config.OFFLINE {
		if config.NETWORK_MODE.Reason == config.NETWORK_IS_DOWN {
			log.Fatal("You need to be connected to the internet in order to sync your offline and online data.")
		}

		if config.NETWORK_MODE.Reason == config.SERVICE_IS_DOWN {
			log.Fatal("Sorry. We can't sync your offline and online data right now because the service is unavailable.")
		}

		if config.NETWORK_MODE.Reason == config.CONFIG_IS_MISSING {
			log.Fatal("You need to have a configuration file set in order to sync your offline and online data.")
		}
	}

	// TODO: Sync records
}
