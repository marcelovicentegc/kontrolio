package cli

import (
	"github.com/marcelovicentegc/kontrolio-cli/db"
)

func punch() {
	db.SaveRecord()
}
