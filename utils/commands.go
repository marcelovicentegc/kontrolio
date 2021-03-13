package utils

import (
	"fmt"

	"github.com/marcelovicentegc/kontrolio-cli/config"
)

// DisplayOnlineMessage handles the online message to be displayed,
// wheter a production message, or a development message.
func DisplayOnlineMessage(appConfig config.Config) {
	if (appConfig.Dev == "true") {
		fmt.Println(DEV_ENVIRONMENT)
	} else {
		fmt.Println(YOURE_ONLINE)
	}
}