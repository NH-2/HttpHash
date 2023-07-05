package main

import (
	"flag"
	"fmt"
	"os"
)

const STATUS_CODE_ERROR_NO_ARGS = 1
const STATUS_CODE_ERROR_REQUESTS_ERROR = 2

func main() {
	flag.Parse()
	urls := flag.Args()
	if len(urls) == 0 {
		err := fmt.Errorf("no URLs provided. Please input at least one URL to run the program")
		fmt.Printf("error: %v\n", err)
		os.Exit(STATUS_CODE_ERROR_NO_ARGS)
	}

	hash, err := hashResponse(urls[0])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(STATUS_CODE_ERROR_REQUESTS_ERROR)
	}
	fmt.Printf("%s  %s\n", urls[0], hash)

}
