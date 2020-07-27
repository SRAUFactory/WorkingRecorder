package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

const logFileName = "./record.log"
const datetimeFormat = "2006-01-02 15:04:05"

func main() {
	flag.Parse()
	args := flag.Args()
	switch args[0] {
	case "start":
		start()
	case "stop":
		stop()
	case "report":
		records, _ := read()
		report(records)
	case "now":
		records, _ := read()
		current(records)
	default:
		fmt.Println("No such command!!")
	}
}

func start() {
	var work string
	const explain = `
Start your task!!
What do you do? : `

	records, err := read()
	if err != nil {
		records = [][]string{}
	}
	last := len(records) - 1
	if last >= 0 && records[last][1] == "" {
		current(records)
		return
	}
	report(records)

	fmt.Print(explain)
	fmt.Scanf("%s", &work)
	t := time.Now()
	fmt.Printf("Start working on '%s' at %s. Good luck!!", work, t.Format(datetimeFormat))
	fmt.Println()
	log := []string{t.Format(datetimeFormat), "", work}
	save(append(records, log))
}

func stop() {
	var next int
	const explain = `
Please select future plans!!
1: Suspend work
2: Continue to work on other tasks
3: Finish today's work
selectd : `

	records, err := read()
	if err != nil {
		log.Fatal("Please command for 'start'!")
	}

	fmt.Print(explain)
	fmt.Scanf("%d", &next)
	fmt.Printf("selected: %d", next)
	fmt.Println()

	last := len(records) - 1
	if last >= 0 && records[last][1] == "" {
		t := time.Now()
		records[last][1] = t.Format(datetimeFormat)
		save(records)
	}

	if next == 2 {
		start()
	} else if next == 3 {
		t := time.Now()
		err = os.Rename(logFileName, logFileName+t.Format("20060102"))
		if err != nil {
			log.Fatal(err)
		}
		report(records)
	}
}

func report(records [][]string) {
	reports := map[string]time.Duration{}
	var works []string
	var total time.Duration
	for i := range records {
		if records[i][1] == "" {
			records[i][1] = time.Now().Format(datetimeFormat)
		}
		duration := calc(records[i][0], records[i][1])
		work := records[i][2]
		_, ok := reports[work]
		if !ok {
			works = append(works, work)
			reports[work] = duration
		} else {
			reports[work] += duration
		}
		total += duration
	}
	sort.Strings(works)

	fmt.Print("Total work time is ")
	fmt.Println(total)
	for j := 0; j < len(works); j++ {
		fmt.Printf("%s :: %s", works[j], reports[works[j]])
		fmt.Println()
	}
}

func current(records [][]string) {
	last := len(records) - 1
	if last < 0 || records[last][1] != "" {
		log.Fatal("Current work is nothing!!")
	}
	fmt.Printf("Current work is %s", records[last][2])
	fmt.Println()
	fmt.Printf("From %s", records[last][0])
	fmt.Println()
	fmt.Printf("Working time is %s", calc(records[last][0], time.Now().Format(datetimeFormat)))
	fmt.Println()
}

func save(records [][]string) {
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		fmt.Println("Open file error")
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.WriteAll(records)
	if err != nil {
		log.Fatal(err)
	}
	writer.Flush()
}

func read() ([][]string, error) {
	file, err := os.Open(logFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

func calc(start string, stop string) time.Duration {
	startTime, _ := time.Parse(datetimeFormat, start)
	stopTime, _ := time.Parse(datetimeFormat, stop)
	return stopTime.Sub(startTime)
}
