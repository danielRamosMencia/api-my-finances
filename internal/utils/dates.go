package utils

import "time"

func GetNowFormat() string {
	return time.Now().UTC().Format(time.RFC3339Nano)
}
