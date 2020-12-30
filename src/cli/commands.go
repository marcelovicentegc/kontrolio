package cli

import (
	"fmt"
	"log"

	"github.com/marcelovicentegc/kontrolio-cli/src/config"
	"github.com/marcelovicentegc/kontrolio-cli/src/db"
)

func punch() {
	db.SaveRecord()

	if config.NETWORK_MODE.Status == config.ONLINE {
		appConfig := config.GetConfig()
		fmt.Println(appConfig.ApiKey)

		// TODO: Hit the server and save data there
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
