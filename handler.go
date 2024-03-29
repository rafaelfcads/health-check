package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	showIntroduction()
	showMenu()
	command := readCommand()

	switch command {
	case 1:
		startMonitor()
	case 2:
		showLog()
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

func startMonitor() {
	fmt.Println("Monitoring ...")
	// site := "https://banco-questoes-api-es-development.p4ed.com/"
	// sites := []string{site}

	sites := readFile()

	for _, site := range sites {
		validateSite(site)
	}
}

func validateSite(site string) {
	const statusCodeSuccess = 200

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Err", err)
	}

	fmt.Println(resp)

	if resp.StatusCode == statusCodeSuccess {
		registreLog(site, resp.StatusCode)
		fmt.Println("Api successfully loaded. Status ", resp.StatusCode)
		return
	}

	registreLog(site, resp.StatusCode)
}

func readFile() []string {
	var sites []string
	// file, err := ioutil.ReadFile("sites.txt")

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Err", err)
	}

	read := bufio.NewReader(file)

	for {

		line, err := read.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)
		fmt.Println("Line: ", line)

		if err == io.EOF {
			break
		}

		fmt.Println("Line: ", line)
	}

	file.Close()

	return sites
}

func registreLog(site string, status int) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Err", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + site + strconv.Itoa(status) + "\n")

	fmt.Println(file)

	file.Close()
}

func showLog() {

}
