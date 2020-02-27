package main

import (
	"fmt"
	"net/http"
	"os"
)

func mirroredQuery() string {
	responses := make(chan string, 4)

	// goroutine leak possible - slower goroutines would have gotten stuck
	// trying to send their responses on a channel from which no goroutine
	// will ever receive
	go func() { responses <- request("https://ogorodniki.com") }()
	go func() { responses <- request("https://memo.ua") }()
	go func() { responses <- request("https://yandex.ru") }()
	go func() { responses <- request("https://google.com") }()

	return <-responses // return the quickest response
}

func request(hostname string) (response string) {
	resp, err := http.Get(hostname)
	defer resp.Body.Close()
	if err != nil {
		os.Exit(1)
	}

	return hostname
}

func main() {
	resp := mirroredQuery()
	fmt.Println(resp)
}
