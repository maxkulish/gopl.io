// Fetch with io.Copy to copy the response body to us.Stdout without buffer
// Exercise 1.7 and 1.8
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main()  {
	for _, url := range os.Args[1:] {

		// http:// or https:// not found
		if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			url = strings.Replace(url, "http://", "https://", 1)
		} else {
			url = "https://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Response code: %v\n", resp.StatusCode)

		written, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil || written == 0 {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}


	}
}


