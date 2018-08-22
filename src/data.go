package quiz

import (
	"time"
)

// Data a containter for entries in log data
type Data struct {
	IP                 string
	StartTime          time.Time
	EndTime            time.Time
	TimeTaken          int
	OriginalTimeFormat string
}

// DataCollection is need to allow the sorting of the data
type DataCollection []Data

func (d DataCollection) Len() int      { return len(d) }
func (d DataCollection) Swap(i, j int) { d[i], d[j] = d[j], d[i] }
func (d DataCollection) Less(i, j int) bool {
	if d[i].StartTime.Before(d[j].StartTime) {
		return true
	}

	if compareTime(d[i].StartTime, d[j].StartTime) {
		if d[i].EndTime.Before(d[j].EndTime) {
			return true
		}
	}

	return false
}
