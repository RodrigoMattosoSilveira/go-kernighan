// Minimal echo server
// Build server
//
//	     ensure you are on the server3 folder: ~/projects/go-kernighan/ch01/07-web-server/server3
//	     Ensure the Makefile APP attribute is set to: server3
//			$ make
//
// Launch server
//
//	     $ ~/projects/go-kernighan/bin/server3 &, or
//			$ make run
//
// Trigger server behavior
//
//	$ fetch :8080
//
// Remove server
//
//	Find out the port server2 is listening to
//		$ sudo lsof -i :<PortNumber>
//
//	Kill the process
//		$ kill -9 <<process id>>
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)      // call handler for all requests starting with `/'
	http.HandleFunc("/count", counter) // call handler for all requests starting with `/count'
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL Path = %q\n", r.URL.Path)
}

// echoes the number of call thus far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}
