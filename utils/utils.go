package utils

import (
	"time"
)

type Record struct {
	Time time.Time
	Type string
}

// BeginningOfDay returns the given time object formatted to its
// first millisecond of its day
func BeginningOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the given time object formatted to its
// last millisecond of its day
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 999999999, t.Location())
}

func SubtractTime(time1, time2 time.Time) int64 {
	return time2.Sub(time1).Nanoseconds()
}

func IndexOf(haystack []Record, needle Record) int {
	for index, value := range haystack {
		if value.Time.Format(time.RFC3339) == needle.Time.Local().Format(time.RFC3339) {
			return index
		}
	}
	return -1
}
