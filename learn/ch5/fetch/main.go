package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns
// the name and length of the local file
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	file, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	n, err = io.Copy(file, resp.Body)
	// Close file, but prefer error from Copy, if any
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

func main() {
	local, n, err := fetch("https://ogorodniki.com/error")
	fmt.Printf("Path: %s\nSize: %d\nError:%v", local, n, err)
}
