package quiz

import (
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

type testData struct {
	IP        string
	EndTime   string
	TimeTaken int
}

func TestActiveConnecions(t *testing.T) {
	tests := []struct {
		file      string
		inputTime string
		expected  []testData
	}{
		{
			filepath.Join("test_data", "log.csv"),
			"2017-10-23T11:58:00.000",
			[]testData{},
		},
		{
			filepath.Join("test_data", "log.csv"),
			"2017-10-23T12:00:00.000",
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
			"2017-10-23T12:22:00.000",
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
		sort.Sort(data)

		//Start of test
		activeConnections, err := data.ActiveConnections(test.inputTime)
		if err != nil {
			t.Fail()
		}
		sort.Sort(activeConnections)

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
		sort.Sort(data)

		priorStartTime := data[0].StartTime
		priorEndTime := data[0].EndTime
		for _, i := range data[1:] {
			if i.StartTime.Before(priorStartTime) {
				t.Errorf("Case %d: Not sorted properly. Run go test -v to logs", index)
				t.Log(i.StartTime, "    ", i.EndTime, "      ", i.OriginalTimeFormat, i.TimeTaken)
			}
			if compareTime(priorStartTime, i.StartTime) {
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
