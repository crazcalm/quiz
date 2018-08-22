package quiz

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
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
		return []Data{}, fmt.Errorf(ImportCSVError, err.Error())
	}
	defer file.Close()

	data := []Data{}
	var skipHeader bool
	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return []Data{}, err
		}
		if !skipHeader {
			skipHeader = true
			continue
		}

		// CSV is missing Data
		if len(record) < 3 {
			return []Data{}, fmt.Errorf(ImportCSVError, "Expected at least 3 rows")
		}

		// Parsing the endTime
		endTime, err := ParseDate(record[1])
		if err != nil {
			return []Data{}, err
		}

		timeTaken, err := strconv.Atoi(record[2])
		if err != nil {
			return []Data{}, err
		}

		startTime := endTime.Add(time.Duration(time.Duration(-timeTaken) * time.Millisecond))
		data = append(data, Data{record[0], startTime, endTime, timeTaken, record[1]})
	}

	return data, nil
}
