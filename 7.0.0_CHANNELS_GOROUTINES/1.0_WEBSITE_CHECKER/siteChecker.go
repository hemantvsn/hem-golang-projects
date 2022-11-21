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
		"http://hemyaasda.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, site := range sites {
		// so that it runs in parallel
		go checkSiteStatus(site)
	}
	var endTime = time.Now()
	fmt.Println("ENDED program at ", endTime)
	fmt.Println("Execution time = ", endTime.Sub(startTime))

	// Sleep the main routine so that it never exits before any child routine
	// Though this can work, we wont know how much time we want it to sleep always.
	// Hence we need a better way ie CHANNELS
	time.Sleep(5 * time.Second)
}

func checkSiteStatus(site string) {

	var resp, err = http.Get(site)
	if nil != err {
		fmt.Println("Site = ", site, " seems down =", err)
	} else {
		fmt.Println("Site = ", site, " seems UP. HTTP_STATUS = ", resp.StatusCode)
	}

}
