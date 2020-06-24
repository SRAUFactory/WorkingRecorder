package main

import (
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
	output("start", work, t)
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
	output("stop", "", time.Now())
	if next == 2 {
		start()
	} else if next == 3 {
		report()
	}
}

func report() {
	fmt.Println("report")
}

func output(prefix string, work string, t time.Time) {
	f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "%s :: %s :: %s", t.Format(datetimeFormat), prefix, work)
}
