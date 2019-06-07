package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"gopl.io/ch5/links"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests
var tokens = make(chan struct{}, 20)

// allowed domains to crawl if not in map - skip
var allowedDomains = map[string]bool{
	"memo.ua":        true,
	"ogorodniki.com": true,
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}

	return list
}

func main() {
	worklist := make(chan []string)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments
	n++
	go func() { worklist <- os.Args[1:] }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {

			// parse url to extract host name
			parse, err := url.Parse(link)
			if err != nil {
				log.Println(err)
			}

			domain := parse.Host

			if !seen[link] && allowedDomains[domain] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
