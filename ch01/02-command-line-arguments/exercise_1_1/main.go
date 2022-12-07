// Modified to print the command name too
package main

import (
	"fmt"
	"os"
	"strings"
)

// Notice the difference between
// go run main.go arg1 arg2
// go build main.go && ./exercise_1_1 arg1 arg2
func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}
