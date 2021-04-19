package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		} else {
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
	}

	for line, num := range counts {
		if num > 1 {
			fmt.Printf("%s\t%d\n", line, num)
		}
	}
}
