// Reads text (lines) from from filename provided as arguments, inj batch mode; when done, for each line with more
// than one instance, prints the number of instances and the text
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		// ioutil.ReadFile deprecated: As of Go 1.16, this function simply calls os.ReadFile.
		//f, err := ioutil.ReadFile(fileName)
		data, err := os.ReadFile(fileName)
		if err != nil {
			// Go compiler does not accept this
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
