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
