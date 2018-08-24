package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"time"
)

var (
	// ImportCSVError Generic error formating for CSV errors
	ImportCSVError = "Import CSV error: %s"
)

// ImportLogsFromCSV retrieves csv data from file
func ImportLogsFromCSV(fileName string) (DataCollection, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return DataCollection{}, fmt.Errorf(ImportCSVError, err.Error())
	}
	defer file.Close()

	data := DataCollection{}
	var skipHeader bool
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return DataCollection{}, err
		}
		if !skipHeader {
			skipHeader = true
			continue
		}

		// CSV is missing Data
		if len(record) < 3 {
			return DataCollection{}, fmt.Errorf(ImportCSVError, "Expected at least 3 rows")
		}

		// Parsing the endTime
		endTime, err := ParseDate(record[1])
		if err != nil {
			return DataCollection{}, err
		}

		timeTaken, err := strconv.Atoi(record[2])
		if err != nil {
			return DataCollection{}, err
		}

		startTime := endTime.Add(time.Duration(time.Duration(-timeTaken) * time.Millisecond))
		data = append(data, Data{record[0], startTime, endTime, timeTaken, record[1]})
	}

	sort.Sort(data)
	return data, nil
}
