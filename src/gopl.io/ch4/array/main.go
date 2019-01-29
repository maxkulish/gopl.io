package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
	UAH
)

func main() {
	symbol := [...]string{
		USD: "$",
		EUR: "€",
		GBP: "£",
		RMB: "¥",
		UAH: "₴",
	}

	r := [...]int{99, -1}

	fmt.Println(symbol)
	fmt.Println(r)
}
