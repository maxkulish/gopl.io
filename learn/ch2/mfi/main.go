package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/lengthconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mfi: %v\n", err)
			os.Exit(1)
		}

		f := lengthconv.Feet(l)
		m := lengthconv.Meter(l)
		i := lengthconv.Inches(l)

		fmt.Printf("%s = %s, %s = %s\n%s = %s, %s = %s\n%s = %s, %s = %s\n",
			m, lengthconv.MToF(m), m, lengthconv.MToI(m),
			f, lengthconv.FToI(f), f, lengthconv.FToM(f),
			i, lengthconv.IToF(i), i, lengthconv.IToM(i),
		)
	}
}
