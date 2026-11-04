package utils

import (
	"errors"
	"time"
)

func GetNowFormat() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}

// GetMonthRange returns the first and last day of the month as "YYYY-MM-DD" (UTC), the same prefix
// that GetNowFormat() dates start with. Compare them with substr(created_at, 1, 10) BETWEEN ? AND ?.
func GetMonthRange(year, month int) (string, string, error) {
	if year < 1970 || year > 9999 || month < 1 || month > 12 {
		return "", "", errors.New("invalid year or month")
	}

	first := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	last := first.AddDate(0, 1, -1)

	return first.Format(time.DateOnly), last.Format(time.DateOnly), nil
}
