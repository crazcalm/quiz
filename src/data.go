package quiz

import (
	"sort"
	"time"
)

// CompareTime checks to see if two times are equal
func CompareTime(x, y time.Time) bool {
	if x.Year() != y.Year() {
		return false
	}

	if x.Month() != y.Month() {
		return false
	}

	if x.Day() != y.Day() {
		return false
	}

	if x.Hour() != y.Hour() {
		return false
	}

	if x.Minute() != y.Minute() {
		return false
	}

	if x.Nanosecond() != y.Nanosecond() {
		return false
	}

	return true
}

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

	if CompareTime(d[i].StartTime, d[j].StartTime) {
		if d[i].EndTime.Before(d[j].EndTime) {
			return true
		}
	}

	return false
}

/*ActiveConnections algorithm:
---------------

Given the collection is sorted by start time and time taken (secondary sort),
Then any start times post input time T can be ignored.

Thus, I can find the first index X of the collection that has a start time
that starts after input time T.

Then I can iterate from X - 1 down and mark every case Y where the end time
is after the input time T.

All of the marked items will consist of all the connections that were in
progress at input Time T.
*/
func (d DataCollection) ActiveConnections(timestamp string) (DataCollection, error) {
	inputTime, err := ParseDate(timestamp)
	if err != nil {
		return []Data{}, err
	}

	return d.activeConnections(inputTime), nil
}

func (d DataCollection) activeConnections(timestamp time.Time) DataCollection {
	results := DataCollection{}
	//startIndex := d.findStartIndex(timestamp) //  buggy

	for i := len(d) - 1; i > 0; i-- {
		// An EndTime that ends before the input time starts
		// means that we have left the range of possible active
		// connections.
		if d[i].EndTime.Before(timestamp) {
			break
		}

		if d[i].StartTime.Before(timestamp) && CompareTime(d[i].EndTime, timestamp) {
			results = append(results, d[i])
			continue // Prvents d[i] from statisfying both if statements and, as a result, be double counted
		}

		if d[i].StartTime.Before(timestamp) && d[i].EndTime.After(timestamp) {
			results = append(results, d[i])
		}
	}
	sort.Sort(results)
	return results
}

/* findStartIndex is used to the best index in which to start the ActiveConnections iterative search.

It does this via a tradtional binary search. However, (due to timestamps not being unique) if the
passed in timestamp is found, we then need to iterate down the DataCollection to the last instance
that time stamp.

If the timestamp is not found, we then will return the left index because the timestamp at that
index will be greater than the timestamp we were looking for.
*/
func (d DataCollection) findStartIndex(timestamp time.Time) int {
	var result int
	left := 0
	right := len(d) - 1

	for left <= right {
		middle := (left + right) / 2
		if CompareTime(d[middle].StartTime, timestamp) {
			result = middle
			break
		}

		if d[middle].StartTime.Before(timestamp) {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	if result == 0 { // Means that the binary search did not find the result
		return left
	}

	// Need to ensure that the found timestamp is the last instance of this timestamp
	for ; result < len(d); result++ {
		if d[result].StartTime.After(timestamp) {
			return result
		}
	}
	return result
}

// StatsData returns Timestamps, which can be used to compute the basic stats
// based on the DataCollection.
func (d DataCollection) StatsData() Timestamps {
	statsData := Timestamps{}
	seen := make(map[string]bool) // Used to filter out timestamps that have already been seen.
	for _, data := range d {
		_, ok := seen[data.EndTime.String()]
		if ok {
			continue
		}
		seen[data.EndTime.String()] = true
		activeConnections := d.activeConnections(data.EndTime)
		timestampInfo := TimestampInfo{data.EndTime, activeConnections, len(activeConnections)}
		statsData = append(statsData, timestampInfo)
	}
	sort.Sort(statsData)
	return statsData
}
