package utils

import (
	"fmt"
	"time"
)

// GetTimeFormat gets time format as 20240408
func GetTimeFormat(t time.Time) string {
	return fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
}
