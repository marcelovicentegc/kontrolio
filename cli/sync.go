package cli

import (
	"log"

	"github.com/marcelovicentegc/kontrolio-cli/client"
	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

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