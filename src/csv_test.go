package quiz

import (
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestImportLogsFromCSV(t *testing.T) {
	nanosecondsToSecond := 1000000000   // Reference https://duckduckgo.com/?q=seconds+to+nano+seconds&t=canonical&ia=answer
	nanosecondsToMillisecond := 1000000 // Reference https://duckduckgo.com/?q=milliseconds+to+nano+seconds&t=canonical&ia=answer

	tests := []struct {
		file         string
		numOfEntries int
		IP           []string
		startTime    []time.Time
		endTime      []time.Time
	}{
		{
			filepath.Join("test_data", "good_data.csv"),
			2,
			[]string{"1.2.3.4", "1.2.3.5"},
			[]time.Time{
				time.Date(2017, time.Month(10), 23, 11, 59, 59, nanosecondsToSecond-20*nanosecondsToMillisecond, new(time.Location)),
				time.Date(2017, time.Month(10), 23, 12, 0, 0, 8, new(time.Location)),
			},
			[]time.Time{
				time.Date(2017, time.Month(10), 23, 12, 0, 0, 0, new(time.Location)),
				time.Date(2017, time.Month(10), 23, 12, 0, 1, 0, new(time.Location)),
			},
		},
	}

	for index, test := range tests {
		data, err := ImportLogsFromCSV(test.file)
		if err != nil {
			t.Errorf("Case %d: Unexpected error: %s", index, err.Error())
			continue
		}

		if len(data) != test.numOfEntries {
			t.Errorf("Case %d: Expected %d entries, but got %d", index, test.numOfEntries, len(data))
			t.Log(data)
		}

		if !strings.EqualFold(data[index].IP, test.IP[index]) {
			t.Errorf("Case %d: Expected %s but got %s", index, test.IP[index], data[index].IP)
		}

		if !CompareTime(data[index].StartTime, test.startTime[index]) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.startTime[index], data[index].StartTime)
		}

		if !CompareTime(data[index].EndTime, test.endTime[index]) {
			t.Errorf("Case %d: Expected %v, but got %v", index, test.endTime[index], data[index].EndTime)
		}

	}
}

func TestImportLogsFromCSVErrors(t *testing.T) {
	tests := []struct {
		file  string
		error string
	}{
		{"badfileName", "Import CSV error: open badfileName: no such file or directory"},
		{filepath.Join("test_data", "missing_data.csv"), "record on line 4: wrong number of fields"},
		{filepath.Join("test_data", "bad_time.csv"), "Expected time in format '2017-10-23T12:20:03.000', but got 12:50:00.000"},
		{filepath.Join("test_data", "bad_timetaken.csv"), "strconv.Atoi: parsing \"bad\": invalid syntax"},
	}

	for index, test := range tests {
		_, err := ImportLogsFromCSV(test.file)
		if err == nil {
			t.Errorf("Case %d: Expected err '%s', but did not get it", index, test.error)
			continue
		}
		if !strings.EqualFold(test.error, err.Error()) {
			t.Errorf("Case %d: Expected '%s', but got '%s'", index, test.error, err.Error())
		}
	}
}
