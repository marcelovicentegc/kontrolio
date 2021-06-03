package cli

import (
	"time"

	"github.com/cheynewallace/tabby"
	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

func plotTable() {
	records := db.GetRecords()
	table := tabby.New()
	table.AddHeader("Timestamp", "Record")

	logBuilder(records, func(workNanoseconds *int64, workWindowNanoseconds *int64, currentDay *time.Time, record *utils.Record) {
		if record != nil {
			table.AddLine(record.Time.Format(time.RFC3339), record.Type)
		}
	})

	table.Print()
}
