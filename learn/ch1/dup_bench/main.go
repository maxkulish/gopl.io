package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	counts := make(map[string]int)

	for _, fileName := range os.Args[1:] {
		ReadWholeFile(fileName, counts)
		ReadFileByLine(fileName, counts)
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d: %s\n", n, line)
		}
	}
}

func ReadWholeFile(fileName string, counts map[string]int) {

	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
	}

	for _, line := range strings.Split(string(bytes), "\n") {
		counts[line]++
	}
}

func ReadFileByLine(fileName string, counts map[string]int) {

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	f.Close()
}
