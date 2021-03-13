package utils

import (
	"time"

	"github.com/kyokomi/emoji/v2"
)

const (
	SYNC_OFFLINE         = "You need to be connected to the internet in order to sync your offline and online data."
	SYNC_SERVICE_DOWN    = "Sorry. We can't sync your offline and online data right now because the service is unavailable."
	SYNC_CONFIG_MISSING  = "You need to have a configuration file set in order to sync your offline and online data."
	SYNC_API_KEY_MISSING = "You need to have an API key set on your configuration file in order to sync your offline and online data. Sign up @ https://kontrolio.com to get an API key."
	WORKDAY_STATUS       = "\nToday, you've worked: "
	PUNCH_SUCCESS        = " sucessfully."
	DEV_ENVIRONMENT = "ATTENTION! You're online but this is a DEVELOPMENT environment."
	CREATING_BUCKET = "\nBucket doesn't exist, creating it...\n"
)

var (
	YOURE_ONLINE  = emoji.Sprint("\n:earth_americas:You're online.\n")
	YOURE_OFFLINE = emoji.Sprint("\n:mobile_phone_off:You're offline.\n")
)

func FormatLogMessageFooter(workTime string, workWindowTime string) string {
	return ColorGreen + "Worked " + workTime + " in a " + workWindowTime + " work window." + ColorReset + "\n"
}

func FormatPunchMessage(recordType string) string {
	return "Punched " + recordType + PUNCH_SUCCESS + "\n"
}

func FormatLogMessageHeader(currentDay *time.Time) string {
	return ColorCyan + "\n" + currentDay.Format(time.RFC850) + ColorReset
}