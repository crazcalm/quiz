package quiz

import (
	"errors"
	"math"
	"time"
)

// TimestampInfo is used to collect all the needed active
// connections data per a specific timestamp
type TimestampInfo struct {
	Timestamp              time.Time
	ActiveConnections      DataCollection
	TotalActiveConnections int
}

// Timestamps is a collection of TimestampInfo that is used
// to provide basics stats, such as min, max, and average number of
// connections.
type Timestamps []TimestampInfo

func (t Timestamps) Len() int      { return len(t) }
func (t Timestamps) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t Timestamps) Less(i, j int) bool {
	return t[i].TotalActiveConnections < t[j].TotalActiveConnections
}

// Min returns the timestamp with the least number of active connections
func (t Timestamps) Min() TimestampInfo {
	if len(t) > 0 {
		return t[0]
	}
	return TimestampInfo{}
}

// Max returns the timestamp with the most numer of active connections
func (t Timestamps) Max() TimestampInfo {
	if len(t) > 0 {
		return t[len(t)-1]
	}
	return TimestampInfo{}
}

// Average returns the numerator and denominator for the fraction used to
// compute the average number of active connections
func (t Timestamps) Average() (int, error) {
	if len(t) == 0 {
		return 0, errors.New("Cannot find average of empty Timestamp collection")
	}
	numerator := 0
	for _, item := range t {
		numerator += item.TotalActiveConnections
	}
	//The returned result is undergoing float division, being rounded, and then converted to an int
	return int(math.Round(float64(numerator) / float64(len(t)))), nil
}
