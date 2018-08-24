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

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Helllloooo!!!!!\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(1)
	}

	for i, arg := range os.Args {
		fmt.Println(i, arg)
	}

	file := os.Args[1]
	inpuTimes := os.Args[2]
	timestamps := strings.Split(inpuTimes, ",")

	times := []time.Time{}
	for _, i := range timestamps {
		tempt, err := cleanAndFormatTime(i)
		if err != nil {
			fmt.Println(err)
		}

		times = append(times, tempt)
	}

	fmt.Println(file)
	fmt.Println(times)

	dataCollection, err := quiz.ImportLogsFromCSV(file)
	if err != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", err.Error())
	}
	fmt.Println(dataCollection)
}
