package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch4/imdb"
)

func main() {
	result, err := imdb.SearchMovie(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", len(result.Movies))
	for _, item := range result.Movies {
		fmt.Printf("Title: %s, year: %s, type: %v\n", item.Title, item.Year, item.Type)
	}

}
