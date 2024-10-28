package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const numPings = 10
const delayValue = 1
const filename = "go-forest/sites.txt"

func main() {
	welcome()

	for {
		menu()
		getFileContent(filename)
		choice := choiceService()

		switch choice {
		case 1:
			fmt.Println("Starting monitoring...")
			// sites := getSitesSlice()
			sites := getSitesFromFile(filename)
			for j := 0; j < numPings; j++ {
				for i := 0; i < len(sites); i++ {
					monitorSite(sites[i])
				}
				fmt.Println("----------------------")
				time.Sleep(delayValue * time.Minute)
			}
		case 2:
			fmt.Println("Logs...")
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("This is not valid!")
			os.Exit(-1)
			// main()
		}
	}
}

func welcome() {
	name := "Carmelita"
	fmt.Println(">>>Hey,", name, "\b. Welcome to MoniHTTPtoring!<<<")
}

func menu () {
	fmt.Println("Here's the provided services:")
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
	fmt.Println("Choose an option:")
}

func choiceService() int {
	var choice int
	fmt.Scan(&choice)

	return choice
}

func monitorSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("An error has occurred:", err)
	}

	statusCode := response.StatusCode
	if statusCode == 200 {
		fmt.Println(site, "website was successfully pinged!")
	} else {
		fmt.Println("Something's wrong with", site, "website. Status code:", statusCode)
	}
}

func getSitesArray() [4]string {
	var sites [4]string
	sites[0] = "https://www.nytimes.com/international/"
	sites[1] = "https://www.zeit.de/index"
	sites[2] = "https://www.lemonde.fr/"
	sites[3] = "https://www.folha.uol.com.br/"

	for i, site := range sites {
		fmt.Println("On", i, "th position, is", site, "website.")
	}

	return sites
}

func getSitesSlice() []string {
	sites := []string{"https://www.hackerrank.com/dashboard", "https://leetcode.com/u/CarmelitaBraga/"}
	sites = append(sites, "https://youtube.com/")
	sites = append(sites, "https://www.louvre.fr/")
	return sites
}

func getSitesFromFile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("An error has occurred", err)
	}

	reader := bufio.NewReader(file)
	sites := []string{}

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}

func getFileContent(filename string) {
	file, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}

	fmt.Println(string(file))
}