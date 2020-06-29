package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
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
		report()
	default:
		fmt.Println("No such command!!")
	}
}

func start() {
	var work string
	fmt.Println("start")
	const explain = `
Start your task!!
What do you do? : `
	fmt.Print(explain)
	fmt.Scanf("%s", &work)
	t := time.Now()
	fmt.Printf("Start working on '%s' at %s. Good luck!!", work, t.Format(datetimeFormat))
	fmt.Println()
	output(work, t)
}

func stop() {
	var next int
	fmt.Print("stop")
	const explain = `
Please select future plans!!
1: Suspend work
2: Continue to work on other tasks
3: Finish today's work
selectd : `
	fmt.Print(explain)
	fmt.Scanf("%d", &next)
	fmt.Printf("selected: %d", next)
	fmt.Println()
	output("", time.Now())
	if next == 2 {
		start()
	} else if next == 3 {
		report()
	}
}

func report() {
	fmt.Println("report")
}

func output(work string, t time.Time) {
	records, err := read()
	if err != nil {
		records = [][]string{}
	}
	last := len(records) - 1
	if last >= 0 && records[last][1] == "" {
		records[last][1] = t.Format(datetimeFormat)
	} else {
		log := []string{t.Format(datetimeFormat), "", work}
		records = append(records, log)
	}
	save(records)
}

func save(records [][]string) {
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		fmt.Println("ファイル作成エラー")
	}
	defer file.Close()

	log.Printf("%#v", records)
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
