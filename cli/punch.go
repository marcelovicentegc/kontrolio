package cli

import (
	"fmt"

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