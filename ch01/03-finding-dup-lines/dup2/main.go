// Reads text (lines) from stdio or from filename provided as arguments, in streaming move; when done, for each line
// with more than one instance, prints the number of instances and the text
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				// Go compiler does not accept this
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			// This is an IntelliJ bug, it works fine
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			// NOTE: that my linker suggested to use the formatting function instead
			//fmt.Println("%d\t%s\n", n, line)
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}
