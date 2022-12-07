package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// NOTE: Use ctrl+z Windows and ctrl+D for Mac and Linux to stop collecting data
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Error(
	for line, n := range counts {
		if n > 1 {
			// NOTE: that my linker suggested to use the formating function instead
			//fmt.Println("%d\t%s\n", n, line)
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
