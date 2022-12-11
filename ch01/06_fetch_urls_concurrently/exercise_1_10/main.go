// Fetches URLs concurrently and reports their times and sizes
//
// Changed the Makefile run task to fetch the same large size many times
// - Observed two aspects
//   - The fetch time drops almost by half after a few fetches
//   - the first fetch is the fastest!
//
// Do I get the same content all the time?
// - I get the same number of bytes everytime
//
// Write results to a data file
// - see data.txt
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	curr_wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise_1-11: %v\n", err)
		os.Exit(1)
	}
	f, err := os.Create(curr_wd + "/../../../output/exercise_1-10.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "exercise_1-11: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a go routine
	}
	for range os.Args[1:] {
		//fmt.Println((<-ch)) // receive from channel
		fmt.Fprintf(f, <-ch)
	}
	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel
		return
	}
	nbytes, err := io.Copy(io.Discard, res.Body)
	_ = res.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // send to channel
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs %7d %s\n", secs, nbytes, url) // send to channel
}
