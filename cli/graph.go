package cli

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/guptarohit/asciigraph"
	"github.com/marcelovicentegc/kontrolio-cli/db"
	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

func plotGraph() {
	records := db.GetRecords()

	var data []float64

	logBuilder(records, func(workNanoseconds *int64, workWindowNanoseconds *int64, currentDay *time.Time, record *utils.Record) {
		if workNanoseconds != nil && workWindowNanoseconds != nil {
			hours := fmt.Sprintf("%02d.%02d", *workNanoseconds/time.Hour.Nanoseconds(), *workNanoseconds/time.Minute.Nanoseconds())
			formatted, err := strconv.ParseFloat(hours, 64)

			if err != nil {
				log.Fatal(err)
			}

			data = append(data, formatted)
		}
	})

	graph := asciigraph.Plot(data, asciigraph.Height(10), asciigraph.Precision(1), asciigraph.Caption("Daily work hours plot"))

	fmt.Println(graph)
}
