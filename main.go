package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/fatih/color"
)

type headerFlags []string

func (h *headerFlags) String() string {
	return strings.Join(*h, ";")
}

func (h *headerFlags) Set(value string) error {
	*h = append(*h, value)
	return nil
}

func main() {
	color.Yellow("You are using gurl v0.1\n\n")

	urlPtr := flag.String("url", "", "URL to call")
	methodPtr := flag.String("X", "GET", "HTTP method")

	var headers headerFlags
	flag.Var(&headers, "H", "HTTP request headers")

	flag.Parse()
	if *urlPtr == "" {
		color.Red("URL not specified")
		os.Exit(1)
	}

	client := &http.Client{}
	req, err := http.NewRequest(*methodPtr, *urlPtr, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range headers {
		parts := strings.Split(v, ":")
		if len(parts) == 2 {
			req.Header.Add(http.CanonicalHeaderKey(parts[0]), strings.Trim(parts[1], " "))
		}
	}

	color.Magenta("REQUEST")
	color.Cyan("- " + *methodPtr + " " + *urlPtr)
	fmt.Println("")

	color.Magenta("REQUEST HEADERS")
	for key, value := range req.Header {
		color.Cyan("- %s: %s\n", key, strings.Join(value, ","))
	}
	fmt.Println("")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()
	color.Green("RESPONSE STATUS")
	color.Cyan("- " + resp.Status)
	fmt.Println("")

	color.Green("RESPONSE HEADERS")
	for key, value := range resp.Header {
		color.Cyan("- %s: %s\n", key, strings.Join(value, ","))
	}
	fmt.Println("")

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red("Could not read response body")
	} else if len(body) > 0 {
		color.Green("BODY")
		color.Cyan(string(body))
	}
}
