package utils

import (
	"fmt"

	"github.com/marcelovicentegc/kontrolio-cli/config"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
)

// DisplayOnlineMessage handles the online message to be displayed,
// wheter a production message, or a development message.
func DisplayOnlineMessage(appConfig config.Config) {
	if (appConfig.Dev == "true") {
		fmt.Println(messages.DEV_ENVIRONMENT )
	} else {
		fmt.Println(messages.YOURE_ONLINE)
	}
}