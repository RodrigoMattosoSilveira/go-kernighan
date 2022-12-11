// Minimal echo server
// Launch
//
//	$ ./server1
//
// Remove
//
//	Find out the port server2 is listening to
//		sudo lsof -i :<PortNumber>
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
