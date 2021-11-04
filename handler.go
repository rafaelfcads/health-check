package main

import (
	"fmt"
	"os"
)

func main() {
	showIntroduction()
	showMenu()
	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring ...")
	case 2:
		fmt.Println("Logging ...")
	case 0:
		fmt.Println("Exiting ...")
		os.Exit(0)
	default:
		fmt.Println("Command not found!")
		os.Exit(-1)
	}
}

func showIntroduction() {
	name := "Rafael"
	var version float32 = 1.1

	fmt.Println("Name: ", name)
	fmt.Println("Version: ", version)
}

func readCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("The command selected was ", command)

	return command
}

func showMenu() {
	fmt.Println("1- Monitoring Start")
	fmt.Println("2- Show Logs")
	fmt.Println("0- Exit")
}
