package main

import (
	"flag"
	"fmt"
	"os"
)

const STATUS_CODE_ERROR_NO_ARGS = 1

var maxParallelRequests = flag.Uint("parallel", 10, "max requests to send in parallel. positive number.")

func worker(id uint, urls <-chan string, done chan<- bool) {
	// get input from channel
	for url := range urls {
		// start working on url
		hash, fullUrl, err := hashResponse(url)
		if err != nil {
			// encountered an error, log and continue
			fmt.Printf("error: %v\n", err)
			continue
		}
		fmt.Printf("%s  %s\n", fullUrl, hash)
	}
	// push to results channel to finish
	done <- true
}

func main() {
	// parse flags and arguments
	flag.Parse()
	urls := flag.Args()

	// check args are non nil
	if len(urls) == 0 {
		err := fmt.Errorf("no URLs provided. Please input at least one URL to run the program")
		fmt.Printf("error: %v\n", err)
		os.Exit(STATUS_CODE_ERROR_NO_ARGS)
	}

	jobs := make(chan string, *maxParallelRequests)
	results := make(chan bool, *maxParallelRequests)

	for w := uint(1); w <= *maxParallelRequests; w++ {
		go worker(w, jobs, results)
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	for d := uint(1); d <= *maxParallelRequests; d++ {
		<-results
	}
}
