package quiz

import (
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type testData struct {
	IP        string
	EndTime   string
	TimeTaken int
}

func TestStatsData(t *testing.T) {
	tests := []struct {
		file             string
		len              int
		times            []string
		totalConnections []int
	}{
		{
			filepath.Join("test_data", "log.csv"),
			7,
			[]string{
				"2017-10-23 12:05:00 +0000 UTC",
				"2017-10-23 12:20:03 +0000 UTC",
				"2017-10-23 12:50:00 +0000 UTC",
				"2017-10-23 12:00:01 +0000 UTC",
				"2017-10-23 12:20:05 +0000 UTC",
				"2017-10-23 12:22:00 +0000 UTC",
				"2017-10-23 12:00:00 +0000 UTC",
			},
			[]int{
				1, 1, 2, 2, 2, 3, 6,
			},
		},
		{
			filepath.Join("test_data", "empty_log.csv"),
			0,
			[]string{},
			[]int{},
		},
	}

	for index, test := range tests {
		data, err := ImportLogsFromCSV(test.file)
		if err != nil {
			t.Errorf("Case %d: Unexpected Error: %s", index, err.Error())
		}

		//Test starts
		timestampsData := data.StatsData()
		if len(timestampsData) != test.len {
			t.Errorf("Case %d: Expected %d number of entries, but got %d", index, test.len, len(timestampsData))
		}

		for i, time := range test.times {
			if !strings.EqualFold(time, timestampsData[i].Timestamp.String()) || test.totalConnections[i] != timestampsData[i].TotalActiveConnections {
				t.Errorf("Case %d: Expected %s to have %d active connections, but got %s and %d", index, time, test.totalConnections[i], timestampsData[i].Timestamp.String(), timestampsData[i].TotalActiveConnections)
			}
		}

	}
}

func TestActiveConnecions(t *testing.T) {
	tests := []struct {
		file      string
		inputTime time.Time
		expected  []testData
	}{
		{
			filepath.Join("test_data", "log.csv"),
			time.Date(2017, time.Month(10), 23, 11, 58, 00, 000, new(time.Location)),
			[]testData{},
		},
		{
			filepath.Join("test_data", "log.csv"),
			time.Date(2017, time.Month(10), 23, 12, 00, 00, 000, new(time.Location)),
			[]testData{

				testData{
					"1.2.3.24",
					"2017-10-23T12:00:00.000",
					200,
				},

				testData{
					"1.2.3.15",
					"2017-10-23T12:00:01.000",
					1200,
				},

				testData{
					"1.2.3.9",
					"2017-10-23T12:00:00.000",
					100,
				},

				testData{
					"1.2.3.19",
					"2017-10-23T12:00:00.000",
					100,
				},

				testData{
					"1.2.3.4",
					"2017-10-23T12:00:00.000",
					20,
				},

				testData{
					"1.2.3.14",
					"2017-10-23T12:00:00.000",
					20,
				},
			},
		},
		{
			filepath.Join("test_data", "log.csv"),
			time.Date(2017, time.Month(10), 23, 12, 22, 00, 000, new(time.Location)),
			[]testData{
				testData{
					"1.2.3.48",
					"2017-10-23T12:22:00.000",
					600,
				},
				testData{
					"1.2.3.18",
					"2017-10-23T12:22:00.000",
					200,
				},
				testData{
					"1.2.3.8",
					"2017-10-23T12:22:00.000",
					200,
				},
			},
		},
	}

	for index, test := range tests {
		data, err := ImportLogsFromCSV(test.file)
		if err != nil {
			t.Errorf("Case %d: Unexpected Error: %s", index, err.Error())
		}

		//Start of test
		activeConnections := data.ActiveConnections(test.inputTime)

		if len(activeConnections) != len(test.expected) {

			t.Log("input time: ", test.inputTime)

			t.Log("Logs:")
			for _, i := range data {
				t.Log(i)
			}

			t.Log("active connections")
			for _, i := range activeConnections {
				t.Log(i)
			}
			t.Errorf("Case %d: Expected %d elements, but got %d", index, len(test.expected), len(activeConnections))
			continue
		}

		for i, conn := range activeConnections {
			if !strings.EqualFold(conn.IP, test.expected[i].IP) {
				t.Errorf("Case %d: Expected index %d to have an IP name of %s, but got %s", index, i, test.expected[i].IP, conn.IP)
			}

			if !strings.EqualFold(conn.OriginalTimeFormat, test.expected[i].EndTime) {
				t.Errorf("Case %d: Expected index %d to have a EndTime string of %s, but got %s", index, i, test.expected[i].EndTime, conn.OriginalTimeFormat)
			}

			if conn.TimeTaken != test.expected[i].TimeTaken {
				t.Errorf("Case %d: Expected index %d to have a TimeTaken value of %d, but got %d", index, i, test.expected[i].TimeTaken, conn.TimeTaken)
			}
		}

	}
}

func TestSortDataCollection(t *testing.T) {
	tests := []struct {
		file string
	}{
		{filepath.Join("test_data", "log.csv")},
	}

	for index, test := range tests {
		data, err := ImportLogsFromCSV(test.file)
		if err != nil {
			t.Errorf("Case %d: Unexpected Error: %s", index, err.Error())
		}

		priorStartTime := data[0].StartTime
		priorEndTime := data[0].EndTime
		for _, i := range data[1:] {
			if i.StartTime.Before(priorStartTime) {
				t.Errorf("Case %d: Not sorted properly. Run go test -v to logs", index)
				t.Log(i.StartTime, "    ", i.EndTime, "      ", i.OriginalTimeFormat, i.TimeTaken)
			}
			if CompareTime(priorStartTime, i.StartTime) {
				if i.EndTime.Before(priorEndTime) {
					t.Errorf("Case %d: Not sorted properly. Run go test -v to logs", index)
					t.Log(i.StartTime, "    ", i.EndTime, "      ", i.OriginalTimeFormat, i.TimeTaken)
				}
			}

			priorStartTime = i.StartTime
			priorEndTime = i.EndTime
		}
	}
}
