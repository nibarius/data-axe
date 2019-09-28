package main

import (
	"errors"
	"strconv"
	"time"
)

func timestampToDate(timestampString string) (string, error) {
	if timestampString == "now" {
		timestampString = strconv.FormatInt(time.Now().Unix(), 10)
	}

	timestamp, err := strconv.ParseInt(timestampString, 10, 64)
	if err != nil {
		return "", errors.New(timestampString + " is not a valid timestamp")
	}

	date, inMilliseconds := timestampIntToDate(timestamp)
	var precision = "seconds"
	if inMilliseconds {
		precision = "milli" + precision
	}
	params := cardParameters{
		"Timestamp in " + precision,
		timestampString,
		"Date",
		date,
	}

	return getHtmlDocument("Timestamp converter", []cardParameters{params})
}

func timestampIntToDate(timestamp int64) (date string, inMilliseconds bool) {
	if timestamp > 3000000000 {
		inMilliseconds = true
		timestamp = timestamp / 1000
	} else {
		inMilliseconds = false
	}
	date = time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	return
}
