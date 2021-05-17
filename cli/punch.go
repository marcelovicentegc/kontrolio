package cli

import (
	"fmt"

	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/messages"
)

func punch() {
	recordType := db.SaveOfflineRecord()
	fmt.Println(messages.FormatPunchMessage(recordType))
}
