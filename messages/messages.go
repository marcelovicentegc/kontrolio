package messages

import (
	"github.com/kyokomi/emoji/v2"
	"github.com/mgutz/ansi"
)

const (
	SyncOffline            = "You need to be connected to the internet in order to sync your offline and online data."
	SyncServiceDown        = "Sorry. We can't sync your offline and online data right now because the service is unavailable."
	SyncConfigMissing      = "You need to have a configuration file set in order to sync your offline and online data."
	SyncAPIKeyMissing      = "You need to have an API key set on your configuration file in order to sync your offline and online data. Sign up @ https://kontrolio.com to get an API key."
	WordayStatus           = "\nToday, you've worked: "
	PunchSuccess           = " sucessfully."
	DevEnvironment         = "ATTENTION! You're online but this is a DEVELOPMENT environment."
	CreatingBucket         = "\nBucket doesn't exist, creating it...\n"
	FailedParsingRequest   = "Something went wrong while parsing the response body.\n"
	DoneConfiguring        = "\n✨ Nice! You can change what you've just set anytime by running "
	KontrolioConfigCommand = "kontrolio config"
)

var (
	ColorReset     = ansi.ColorFunc("default")
	ColorRed       = ansi.ColorFunc("red")
	ColorGreen     = ansi.ColorFunc("green")
	ColorYellow    = ansi.ColorFunc("yellow")
	ColorBlue      = ansi.ColorFunc("blue")
	ColorMagenta   = ansi.ColorFunc("magenta")
	ColorCyan      = ansi.ColorFunc("cyan")
	ColorWhite     = ansi.ColorFunc("white")
	ColorWhiteBold = ansi.ColorFunc("white+b")
	ColorGreenBold = ansi.ColorFunc("green+b")
)

var (
	IsOnline  = emoji.Sprint("\n:earth_americas:You're online.\n")
	IsOffline = emoji.Sprint("\n:mobile_phone_off:You're offline.\n")
)
