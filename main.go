package main

import (
	"flag"
	"fmt"
)

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
	fmt.Println("start")
}

func stop() {
	var next int
	fmt.Print("stop")
	const explain = `
Please select future plans!!
1: Suspend work
2: Continue to work on other tasks
3: Finish today's work
select: `
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
