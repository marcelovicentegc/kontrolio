package cli

import (
	"fmt"

	"github.com/marcelovicentegc/kontrolio-cli/client"
	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
)

func punch() {
	if config.Network.Status == config.Offline {
		fmt.Println(messages.IsOffline)
		recordType := db.SaveOfflineRecord()
		fmt.Println(messages.FormatPunchMessage(recordType))
		return
	}

	if config.Network.Status == config.Online {
		appConfig := config.GetConfig()
		messages.DisplayOnlineMessage(*appConfig)

		// TODO: Sync remote and local data, then save record remotely and locally.

		recordType := client.CreateRecord(appConfig.ApiKey)
		fmt.Println(messages.FormatPunchMessage(recordType))
	}
}
