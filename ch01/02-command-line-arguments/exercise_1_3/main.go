// Compare for and Join
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep += " "
	}
	fmt.Println(s)
	fmt.Printf("For time: %s\n", time.Since(start))

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Join time: %s\n", time.Since(start))
}
