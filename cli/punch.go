package cli

import (
	"fmt"

	"github.com/marcelovicentegc/kontrolio-cli/client"
	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

func punch() {
	if config.Network.Status == config.Offline {
		fmt.Println(utils.YOURE_OFFLINE)
		recordType := db.SaveOfflineRecord()
		fmt.Println(utils.FormatPunchMessage(recordType))
		return
	}

	if config.Network.Status == config.Online {
		appConfig := config.GetConfig()
		utils.DisplayOnlineMessage(*appConfig)

		// TODO: Sync remote and local data, then save record remotely and locally.

		recordType := client.CreateRecord(appConfig.ApiKey)
		fmt.Println(utils.FormatPunchMessage(recordType))
	}
}