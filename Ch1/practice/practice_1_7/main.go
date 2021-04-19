package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, ioErr := io.Copy(os.Stdout, resp.Body)

		if ioErr != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", arg, ioErr)
			os.Exit(1)
		}
	}
}
