package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"
)

func main() {
	color.Yellow("You are using gurl v0.1\n\n")

	urlPtr := flag.String("url", "", "URL to call")

	flag.Parse()
	if *urlPtr == "" {
		color.Red("URL not specified")
		os.Exit(1)
	}

	resp, err := http.Get(*urlPtr)
	if err != nil {
		color.Red("Request to '%s' failed\n", *urlPtr)
		os.Exit(1)
	}

	defer resp.Body.Close()
	color.Green("STATUS")
	color.Blue("- " + resp.Status)
	fmt.Println("")

	color.Green("HEADERS")
	for key, value := range resp.Header {
		color.Blue("- %s: %s\n", key, value)
	}
	fmt.Println("")

	color.Green("BODY")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red("Could not read response body")
	} else {
		color.Blue(string(body))
	}
}
