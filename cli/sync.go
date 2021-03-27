package cli

import (
	"log"

	"github.com/marcelovicentegc/kontrolio-cli/client"
	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

func sync() {
	if config.Network.Status == config.Offline {
		if config.Network.Reason == config.NetworkIsDown {
			log.Fatal(messages.SyncOffline)
		}

		if config.Network.Reason == config.ServiceIsDown {
			log.Fatal(messages.SyncServiceDown)
		}

		if config.Network.Reason == config.ConfigIsMissing {
			log.Fatal(messages.SyncConfigMissing)
		}

		if config.Network.Reason == config.APIKeyIsMissing {
			log.Fatal(messages.SyncAPIKeyMissing)
		}
	}

	appConfig := config.GetConfig()

	messages.DisplayOnlineMessage(*appConfig)

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
