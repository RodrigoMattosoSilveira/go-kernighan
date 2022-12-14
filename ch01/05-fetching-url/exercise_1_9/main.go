// Write the content at each specified URL, append TOKEN if necessary; does reads and sends to
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const TOKEN = "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(strings.ToLower(url), TOKEN) {
			temp := []string{TOKEN, url}
			url = strings.Join(temp, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, "Response status: %s\n\n", resp.Status)
		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
	}
}
