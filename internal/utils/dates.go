package utils

import (
	"errors"
	"time"
	_ "time/tzdata" // embeds the timezone database: the server image may not have /usr/share/zoneinfo
)

const SQL_DATETIME_FORMAT = "2006-01-02T15:04:05"

func GetNowFormat() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}

type MonthRange struct {
	FirstDay string // "YYYY-MM-DD" in the requested timezone
	LastDay  string
	From     string // first second of the month, in UTC SQL_DATETIME_FORMAT
	To       string // last second of the month, in UTC SQL_DATETIME_FORMAT
}

// GetMonthRange returns the month in the given timezone converted to UTC. Compare From/To
// with substr(created_at, 1, 19) BETWEEN ? AND ?: the seconds prefix of GetNowFormat() dates.
func GetMonthRange(year, month int, location *time.Location) (MonthRange, error) {
	if year < 1970 || year > 9999 || month < 1 || month > 12 {
		return MonthRange{}, errors.New("invalid year or month")
	}

	first := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, location)
	next := first.AddDate(0, 1, 0)

	return MonthRange{
		FirstDay: first.Format(time.DateOnly),
		LastDay:  next.AddDate(0, 0, -1).Format(time.DateOnly),
		From:     first.UTC().Format(SQL_DATETIME_FORMAT),
		To:       next.Add(-time.Second).UTC().Format(SQL_DATETIME_FORMAT),
	}, nil
}
