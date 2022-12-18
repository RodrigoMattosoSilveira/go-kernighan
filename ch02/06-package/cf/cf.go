// cf converts it numeric argument into Celsius and Fahrenheit
package main

import (
	"fmt"
	tempconv "github.com/RodrigoMattosoSilveira/go-kernighan/ch02/06-package/tempcov"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%g = %g, %g = %g\n", f, tempconv.FtoC(f), c, tempconv.CtoF(c))
	}
}
