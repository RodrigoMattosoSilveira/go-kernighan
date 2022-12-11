// Minimal echo server
// Build server
//
//	     ensure you are on the server1 folder: ~/projects/go-kernighan/ch01/07-web-server/server1
//	     Ensure the Makefile APP attribute is set: APP=server1
//			$ make
//
// Launch server
//
//	     $ ~/projects/go-kernighan/bin/server3 &, or
//			$ make run
//
// Trigger server behavior
//
//	$ ~/projects/go-kernighan/bin/exercise_1_8 :8080, or
//  $ make client -Dparm=:8080
//
// Remove server
//
//	Find out the port server2 is listening to
//		$ sudo lsof -i :8080
//
//	Kill the process
//		$ kill -9 <<process id>>

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request call handler
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// handler echoes the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL Path = %q\n", r.URL.Path)
}
