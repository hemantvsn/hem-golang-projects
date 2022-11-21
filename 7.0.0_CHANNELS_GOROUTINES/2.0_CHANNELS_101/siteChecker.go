package main

import (
	"fmt"
	"net/http"
	"time"
)

// Checks the status of given websites to see if they are UP
func main() {
	var startTime = time.Now()

	fmt.Println("Started program at ", startTime)
	var sites = []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}

	// Making a channel of value type string
	var ch = make(chan string)

	for _, site := range sites {
		// so that it runs in parallel
		go checkSiteStatus(site, ch)
	}

	// infinite for loop
	for link := range ch {
		// go routine will send the original link in the channel
		// and we are spawning new goroutine with the same link
		// So this will run indeinitely
		go checkSiteStatus(link, ch)
	}

	var endTime = time.Now()
	fmt.Println("ENDED program at ", endTime)
	fmt.Println("Execution time = ", endTime.Sub(startTime))
}

// ch chan string - ch is var of channel of type string
func checkSiteStatus(site string, ch chan string) {

	var resp, err = http.Get(site)
	if nil != err {
		fmt.Println("Site = ", site, " seems down =", err)
	} else {
		fmt.Println("Site = ", site, " seems UP. HTTP_STATUS = ", resp.StatusCode)
	}

	time.Sleep(2 * time.Second)
	ch <- site

}
