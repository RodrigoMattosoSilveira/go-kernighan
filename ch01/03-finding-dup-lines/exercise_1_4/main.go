// Reads text (lines) from stdio or from filename provided as arguments, in streaming move; when done, it prints all
// the filenames in which each duplicate line occurs
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	//files := make(map[string]string)
	files := make(map[string]string)
	fileNames := os.Args[1:]
	if len(fileNames) == 0 {
		countLines(os.Stdin, counts, files)
	} else {
		for _, arg := range fileNames {
			f, err := os.Open(arg)
			if err != nil {
				// Go compiler does not accept this
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, files)
			// This is an IntelliJ bug, it works fine
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			// NOTE: that my linker suggested to use the formatting function instead
			//fmt.Println("%d\t%s\n", n, line)
			fmt.Printf("%d\t%s\t%s\n", n, line, files[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, files map[string]string) {
	filename := f.Name()
	input := bufio.NewScanner(f)
	for input.Scan() {
		inputText := input.Text()
		counts[inputText]++
		if files[inputText] == "" {
			files[inputText] = filename
		} else {
			if !strings.Contains(files[inputText], filename) {
				files[inputText] = fmt.Sprintf("%s, %s", files[inputText], filename)
			}
		}
	}
}
