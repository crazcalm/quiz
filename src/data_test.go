package quiz

import (
	"path/filepath"
	"sort"
	"testing"
)

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
