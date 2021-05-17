package messages

import (
	"time"

	"github.com/marcelovicentegc/kontrolio-cli/utils"
)

func FormatLogMessageFooter(workTime string, workWindowTime string) string {
	return ColorGreenBold("Worked " + workTime + " in a " + workWindowTime + " work window. \n")
}

func FormatPunchMessage(recordType string) string {
	return "Punched " + recordType + PunchSuccess + "\n"
}

func FormatLogMessageHeader(currentDay *time.Time) string {
	return ColorWhiteBold("\n" + currentDay.Format(time.RFC1123Z)[0:16] + "\n")
}

func FormatLogMessage(record utils.Record) string {
	return record.Time.Format(time.Kitchen) + " " + record.Type + "\n"
}

func FormatStatusMessage(workTimeNanoseconds int64) string {
	return ColorGreenBold(WordayStatus + time.Duration(workTimeNanoseconds).String() + "\n")
}
