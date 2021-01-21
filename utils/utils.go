package utils

import (
	"time"
)

func BeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func SubtractTime(time1, time2 time.Time) int64 {
	return time2.Sub(time1).Nanoseconds()
}

func IndexOf(haystack []string, needle string) int {
	for index, value := range haystack {
		if value == needle {
			return index
		}
	}
	return -1
}

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)