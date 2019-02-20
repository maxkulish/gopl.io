package main

import (
	"fmt"
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// lots of work
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d", x, result) }()
	return x + x
}

func main() {
	bigSlowOperation()
	double(5)
}
