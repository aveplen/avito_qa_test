package utils

import (
	"fmt"
	"strconv"
	"time"
)

func TimestampToTime(timestamp string) (time.Time, error) {
	timstampAsInt, err := strconv.ParseFloat(timestamp, 64)
	if err != nil {
		return time.Now(), fmt.Errorf("can't parse int from timestamp: %w", err)
	}
	return time.Unix(int64(timstampAsInt), 0), nil
}
