// Fetch with io.Copy to copy the response body to us.Stdout without buffer
// Exercise 1.7 and 1.8
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1: %v\n", err)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch read: %v\n", err)
		}
	}
}
