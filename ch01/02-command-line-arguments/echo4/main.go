// prints its command line arguments, very effectively, not not pretty!
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[1:])
}
