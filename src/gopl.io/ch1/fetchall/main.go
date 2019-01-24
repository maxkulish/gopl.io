// Fetchall fetches URLs in parallel and reports their times and sizes.
// Exercise 1.10 - print output to a file
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main()  {
	start := time.Now()
	ch := make(chan string)

	f, err := os.Create("fetchall.txt")
	w := bufio.NewWriter(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't create a file: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch)	// start a goroutine
	}

	for range os.Args[1:] {
		fmt.Fprintf(w, "%v\n", <-ch)
		//fmt.Println(<-ch)	// receive from channel ch
	}

	fmt.Fprintf(w, "%.2fs elapsed\n=====================\n", time.Since(start).Seconds())
	w.Flush()
}

func fetch(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)	// send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %dKb %s", secs, nbytes/1000, url)
}
