package main

import (
	"fmt"
)

func main() {
	var command string
	fmt.Scanf("%s", &command)
	switch command {
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
	fmt.Println("stop")
}

func report() {
	fmt.Println("report")
}
