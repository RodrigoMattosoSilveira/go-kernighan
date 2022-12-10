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
	"strings"
	"time"
)

func main() {
	const TOKEN = "https://"
	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(strings.ToLower(url), TOKEN) {
			temp := []string{TOKEN, url}
			url = strings.Join(temp, "")
		}
		go fetch(url, ch) // start a go routine
	}
	fmt.Fprintf(f, "%s\t\t\t%s\t\t%s\n", "secs", "nbytes", "url") // send to channels
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
	ch <- fmt.Sprintf("%2fs\t\t%7d\t\t%s\n", secs, nbytes, url) // send to channel
}
