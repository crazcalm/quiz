package quiz

import (
	"path/filepath"
	"testing"
)

func TestTimestampsMin(t *testing.T) {
	tests := []struct {
		file     string
		expected int
	}{
		//TODO: test empty log file case
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
		//TODO: test empty log file case
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

func TestTimestampsAverage(t *testing.T) {
	tests := []struct {
		file     string
		expected int
	}{
		//TODO: test empty log file case
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
			t.Errorf("I need to deal with this error: %s", err.Error())
		}

		if test.expected != result {
			t.Errorf("Case %d: Expected %d, but got %d", index, test.expected, result)
		}
	}
}
