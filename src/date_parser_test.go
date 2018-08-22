package quiz

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		timeString  string
		year        int
		month       time.Month
		day         int
		hours       int
		minutes     int
		seconds     int
		nanoSeconds int
	}{
		{
			"2017-10-23T12:15:22.668",
			2017,
			time.Month(10),
			23,
			12,
			15,
			22,
			668,
		},
	}

	for index, test := range tests {
		result, err := ParseDate(test.timeString)
		if err != nil {
			t.Errorf("Case %d: Got an unexpected err: %s", index, err.Error())
			continue
		}

		if result.Year() != test.year {
			t.Errorf("Case %d: Expected year %d, but got %d", index, test.year, result.Year())
		}

		if result.Month() != test.month {
			t.Errorf("Case %d: Expected month %v, but got %v", index, test.month, result.Month())
		}

		if result.Day() != test.day {
			t.Errorf("Case %d: Expected day %d, but got %d", index, test.day, result.Day())
		}

		if result.Hour() != test.hours {
			t.Errorf("Case %d: Expected %d hour(s), but got %d", index, test.hours, result.Hour())
		}

		if result.Minute() != test.minutes {
			t.Errorf("Case %d: Expected %d minute(s), but got %d", index, test.minutes, result.Minute())
		}

		if result.Second() != test.seconds {
			t.Errorf("Case %d: Expected %d second(s), but got %d", index, test.seconds, result.Second())
		}

		if result.Nanosecond() != test.nanoSeconds {
			t.Errorf("Case %d: Expected %d nano second(s), but got %d", index, test.nanoSeconds, result.Nanosecond())
		}
	}
}
