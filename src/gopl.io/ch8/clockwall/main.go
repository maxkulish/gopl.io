// Client of several clock servers at once
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	if len(os.Args) == 1 {
		_, _ = fmt.Fprintln(os.Stderr, "usage: ./clockwall City=host:port ...")
		os.Exit(1)
	}

	flag.Parse()
	serversInput := flag.Args()

	var servers = make(map[string]string)
	for _, server := range serversInput {
		split := strings.Split(server, "=")
		if len(split) != 2 {
			_, _ = fmt.Fprintln(os.Stderr, "bad args: %s\n", split)
			os.Exit(1)
		}

		servers[split[0]] = split[1]
	}

	for city, server := range servers {
		conn, err := net.Dial("tcp", server)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		go printTime(city, os.Stdout, conn)
	}

	// Sleep while other goroutines do the work
	for {
		time.Sleep(time.Minute)
	}
}

func printTime(city string, w io.Writer, r io.Reader) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		_, _ = fmt.Fprintf(w, "%s: %s\n", city, s.Text())
	}
	fmt.Println(city, "done")

	if s.Err() != nil {
		log.Printf("can't read from %s: %s", city, s.Err())
	}
}
