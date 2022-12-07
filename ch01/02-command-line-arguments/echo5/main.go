// prints its command line arguments, iterating over a range
package main

import "strings"

func ForLoop(args []string) string {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep += " "
	}
	//fmt.Println(s)
	return s
}

func JoinLoop(args []string) string {
	//fmt.Println(strings.Join(args[1:], " "))
	s := strings.Join(args[1:], " ")
	return s
}
