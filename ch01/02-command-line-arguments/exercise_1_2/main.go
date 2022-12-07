// prints its command line arguments, iterating over a range and printing the argument order too
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for i := 1; i < len(os.Args); i++ {
		s += sep + "(" + strconv.Itoa(i) + ", " + os.Args[i] + ")"
		sep += "; "
	}
	fmt.Println(s)
}
