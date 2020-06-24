package main

import (
	"flag"
	"fmt"
	"time"
)

const logFileName = "record.log"
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
	// @ToDo 作業記録保存
	if next == 2 {
		start()
	} else if next == 3 {
		report()
	}
}

func report() {
	fmt.Println("report")
}
