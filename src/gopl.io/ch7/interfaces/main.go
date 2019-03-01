package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)

	var any interface{}
	any = true
	any = 12.32
	any = "hello"
	any = map[string]int{"one": 1}
}
