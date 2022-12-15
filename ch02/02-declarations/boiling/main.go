// Package boiling prints the water's boling point
package main

import "fmt"

const BoilingF = 212.0

func main() {
	var f = BoilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("water's boiling poing %g°F or %g°C\n", f, c)
}
