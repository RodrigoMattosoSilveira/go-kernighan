// prints its command line arguments, iterating over a range
package main

import (
	"fmt"
	"strings"
)

func ForLoop(args []string) bool {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep += " "
	}
	fmt.Println(s)
	return true
}

func JoinLoop(args []string) {
	fmt.Println(strings.Join(args[1:], " "))
}
