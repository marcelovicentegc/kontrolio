package messages

import (
	"github.com/mgutz/ansi"
)

const (
	WordayStatus           = "\nToday, you've worked: "
	PunchSuccess           = " sucessfully."
	DevEnvironment         = "ATTENTION! You're online but this is a DEVELOPMENT environment."
	CreatingBucket         = "\nBucket doesn't exist, creating it...\n"
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
