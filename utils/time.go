package utils

import (
	"fmt"
	"time"
)

// GetTimeFormat gets time format as 20240408
func GetTimeFormat(t time.Time) string {
	return fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
}

func SkipWeekends(t time.Time) time.Time {
	// check if t is weekend
	if t.Weekday() == time.Saturday {
		//	add 2 days
		return t.AddDate(0, 0, 2)
	} else if t.Weekday() == time.Sunday {
		//	add 1 day
		return t.AddDate(0, 0, 1)
	}

	return t
}
