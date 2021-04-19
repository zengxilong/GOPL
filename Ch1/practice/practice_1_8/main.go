package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "http://") {
			arg = "http://" + arg
		}
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", arg, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
