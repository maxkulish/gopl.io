// Clock1 is a TCP server that periodically writes the time
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var port = flag.String("port", "8000", "-port 8010")

	flag.Parse()

	adress := fmt.Sprintf("localhost:%s", *port)
	fmt.Println(adress)

	listen, err := net.Listen("tcp", adress)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err) // connection aborted
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
