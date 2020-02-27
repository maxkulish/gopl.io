package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	ch <- "a"
	ch <- "b"
	ch <- "c"

	fmt.Println(<-ch)
	fmt.Println(len(ch))

	fmt.Println(<-ch)
	fmt.Println(len(ch))

	fmt.Println(<-ch)
	fmt.Println(len(ch))

}
