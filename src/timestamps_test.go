package quiz

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestTimestampsMin(t *testing.T) {
	tests := []struct {
		file     string
		expected int
	}{
		{filepath.Join("test_data", "empty_log.csv"), 0},
		{filepath.Join("test_data", "log.csv"), 1},
	}

	for index, test := range tests {
		data, err := ImportLogsFromCSV(test.file)
		if err != nil {
			t.Errorf("Case %d: Unexpected Error: %s", index, err.Error())
		}

		//Test starts
		timestampsData := data.StatsData()

		if test.expected != timestampsData.Min().TotalActiveConnections {
			t.Errorf("Case %d: Expected %d, but got %d", index, test.expected, timestampsData.Min().TotalActiveConnections)
		}
	}
}

func TestTimestampsMax(t *testing.T) {
	tests := []struct {
		file     string
		expected int
	}{
		{filepath.Join("test_data", "empty_log.csv"), 0},
		{filepath.Join("test_data", "log.csv"), 6},
	}

	for index, test := range tests {
		data, err := ImportLogsFromCSV(test.file)
		if err != nil {
			t.Errorf("Case %d: Unexpected Error: %s", index, err.Error())
		}

		//Test starts
		timestampsData := data.StatsData()

		if test.expected != timestampsData.Max().TotalActiveConnections {
			t.Errorf("Case %d: Expected %d, but got %d", index, test.expected, timestampsData.Max().TotalActiveConnections)
		}
	}
}

func TestTimestampsAverageErrors(t *testing.T) {
	tests := []struct {
		file string
		err  string
	}{
		{
			filepath.Join("test_data", "empty_log.csv"),
			"Cannot find average of empty Timestamp collection",
		},
	}

	for index, test := range tests {
		data, err := ImportLogsFromCSV(test.file)
		if err != nil {
			t.Errorf("Case %d: Unexpected Error: %s", index, err.Error())
		}

		//Test starts
		timestampsData := data.StatsData()
		_, err = timestampsData.Average()
		if err == nil {
			t.Errorf("Case %d: Expected an error, but none was found", index)
		}

		if !strings.EqualFold(err.Error(), test.err) {
			t.Errorf("Case %d: Expected '%s', but got '%s'", index, test.err, err.Error())
		}
	}
}

func TestTimestampsAverage(t *testing.T) {
	tests := []struct {
		file     string
		expected int
	}{
		{filepath.Join("test_data", "log.csv"), 2},
	}

	for index, test := range tests {
		data, err := ImportLogsFromCSV(test.file)
		if err != nil {
			t.Errorf("Case %d: Unexpected Error: %s", index, err.Error())
		}

		//Test starts
		timestampsData := data.StatsData()
		result, err := timestampsData.Average()
		if err != nil {
			t.Errorf("Case %d: Unexpected error: %s", index, err.Error())
		}

		if test.expected != result {
			t.Errorf("Case %d: Expected %d, but got %d", index, test.expected, result)
		}
	}
}
