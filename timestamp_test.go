package dataaxe

import (
	"testing"
)

func TestTimestampIntToDate(t *testing.T) {
	var tsInSeconds int64 = 1323782116
	var tsInMilliseconds int64 = 1323782116000
	var expected = "2011-12-13 14:15:16"
	seconds, _ := timestampIntToDate(tsInSeconds)
	milliseconds, _ := timestampIntToDate(tsInMilliseconds)
	if seconds != "2011-12-13 14:15:16" {
		t.Errorf("Timestamp conversion failed. Expected " + expected + " got " + seconds)
	}
	if milliseconds != "2011-12-13 14:15:16" {
		t.Errorf("Timestamp conversion failed. Expected " + expected + " got " + milliseconds)
	}
}
