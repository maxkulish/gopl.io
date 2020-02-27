// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main()  {
	http.HandleFunc("/", handler)	// each request calls handler
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

// handler echoes the HTTP request
func handler(w http.ResponseWriter, r *http.Request)  {
	lissajous(w)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}