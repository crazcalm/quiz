package quiz

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ParseDate is used to parse dates of the format "2017-10-23T12:20:03.000"
// and convert them to a format that go can read.
func ParseDate(timeString string) (time.Time, error) {
	dateTimeParts := strings.Split(timeString, "T")
	if len(dateTimeParts) != 2 {
		return time.Time{}, fmt.Errorf("Expected time in format '%s', but got %s", "2017-10-23T12:20:03.000", timeString)
	}

	dateParts := strings.Split(dateTimeParts[0], "-")
	if len(dateParts) != 3 {
		return time.Time{}, fmt.Errorf("Expected time in format '%s', but got %s", "2017-10-23T12:20:03.000", timeString)
	}

	timeParts := strings.Split(dateTimeParts[1], ":")
	if len(timeParts) != 3 {
		return time.Time{}, fmt.Errorf("Expected time in format '%s', but got %s", "2017-10-23T12:20:03.000", timeString)
	}

	year, err := strconv.Atoi(dateParts[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("Parsing time error: %s", err.Error())
	}

	month, err := strconv.Atoi(dateParts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("Parsing time error: %s", err.Error())
	}
	day, err := strconv.Atoi(dateParts[2])
	if err != nil {
		return time.Time{}, fmt.Errorf("Parsing time error: %s", err.Error())
	}

	hours, err := strconv.Atoi(timeParts[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("Parsing time error: %s", err.Error())
	}

	minutes, err := strconv.Atoi(timeParts[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("Parsing time error: %s", err.Error())
	}

	if len(strings.Split(timeParts[2], ".")) != 2 {
		return time.Time{}, fmt.Errorf("Expected time in format '%s', but got %s", "2017-10-23T12:20:03.000", timeString)
	}

	seconds, err := strconv.Atoi(strings.Split(timeParts[2], ".")[0])
	if err != nil {
		return time.Time{}, fmt.Errorf("Parsing time error: %s", err.Error())
	}

	nanoSeconds, err := strconv.Atoi(strings.Split(timeParts[2], ".")[1])
	if err != nil {
		return time.Time{}, fmt.Errorf("Parsing time error: %s", err.Error())
	}

	return time.Date(year, time.Month(month), day, hours, minutes, seconds, nanoSeconds, new(time.Location)), nil
}
