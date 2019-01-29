package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {

	shaFlag := flag.Int("sha", 256, "someText --sha=512")

	flag.Parse()

	args := flag.Args()

	var result string

	for _, arg := range args {

		switch *shaFlag {
		case 384:
			result += fmt.Sprintf("%x", sha512.Sum384([]byte(arg)))
		case 512:
			result += fmt.Sprintf("%x", sha512.Sum512([]byte(arg)))
		default:
			result += fmt.Sprintf("%x", sha256.Sum256([]byte(arg)))
		}
	}

	fmt.Println(result)
}
