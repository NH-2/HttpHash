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

const (
	usage = `usage: %s -parallel <maxParallelRequests> [URL(s)]...
Makes parallel http requests and prints the address of the request along with the MD5 hash of the response.

example:
%s -parallel 3 example.com google.com reddit.com/r/funny
`
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	// parse flags and arguments
	urls := flag.Args()
	flag.Parse()
	// check args are non nil
	if len(urls) == 0 {
		flag.Usage()
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
