package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/crazcalm/quiz/src"
)

var stats = flag.Bool("stats", false, "Will show the Min, Max and Average number of active connections")

func cleanAndFormatTime(t string) (time.Time, error) {
	t = strings.TrimSpace(t)
	t = strings.Replace(t, " ", "T", 1)
	result, err := quiz.ParseDate(t)

	if err != nil {
		return time.Time{}, errors.New("Check that the timestamp is of the form '2018-10-24 12:13:06.000'")
	}

	return result, nil
}

func hasEnoughArguementsCheck(statsFlag bool) {
	if statsFlag {
		if len(os.Args) < 4 && len(os.Args) != 3 {
			flag.Usage()
			os.Exit(1)
		}

	} else if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(1)
	}
}

func gatherTimes(timestamps []string) (times []time.Time) {
	for _, i := range timestamps {
		tempt, err := cleanAndFormatTime(i)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(2)
		}

		times = append(times, tempt)
	}
	return times
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Helllloooo!!!!!\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Println(os.Args)

	hasEnoughArguementsCheck(*stats)

	//Arguements
	var file string
	var inputTimes string

	// Given that go flags have to go before non-flag args,
	// I need to account for where the file and input times
	// will be. I also need to know if input times are passed in.
	if *stats {
		if len(os.Args) == 3 {
			file = os.Args[len(os.Args)-1]
		} else {
			file = os.Args[len(os.Args)-2]
			inputTimes = os.Args[len(os.Args)-1]
		}
	}

	//Makes sure the flag is not the last arg passed in because Go will not regonize it as a flag
	fmt.Println(strings.TrimSpace(os.Args[len(os.Args)-1]))
	if strings.EqualFold(strings.TrimSpace(os.Args[len(os.Args)-1]), "--stats") {
		fmt.Fprintf(os.Stderr, "Try one of these:\n%s --stats file\n%s --stats file timestamp(s)\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}

	dataCollection, err := quiz.ImportLogsFromCSV(file)
	if err != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", err.Error())
	}

	if !strings.EqualFold(inputTimes, "") {
		timestamps := strings.Split(inputTimes, ",")

		//User input times
		times := gatherTimes(timestamps)

		//Case: Only one input time
		if len(times) == 1 {
			fmt.Println(len(dataCollection.ActiveConnections(times[0])))
		} else { // Muiltiple times
			for _, t := range times {
				fmt.Printf("%s: %d\n", t.Format(time.StampMilli), len(dataCollection.ActiveConnections(t)))
			}
		}

		if *stats {
			//Print a line space for the stats information
			fmt.Println()
		}

	}

	if *stats {
		fmt.Println("Stats")
		statsData := dataCollection.StatsData()
		min := statsData.Min()
		max := statsData.Max()
		fmt.Printf("- Min: %s: %d\n", min.Timestamp.Format(time.StampMilli), min.TotalActiveConnections)
		fmt.Printf("- Max: %s: %d\n", max.Timestamp.Format(time.StampMilli), max.TotalActiveConnections)
		average, err := statsData.Average()
		if err != nil {
			fmt.Println("- Average: 0")
		} else {
			fmt.Printf("- Average: %d\n", average)
		}
	}
}
