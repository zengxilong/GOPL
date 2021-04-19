package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		return
	} else {
		for _, arg := range files {
			counts := make(map[string]int)
			file, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}

			if countLine(file, counts) {
				fmt.Printf("%s\n", arg)
			}
			file.Close()
		}
	}

}

func countLine(f *os.File, counts map[string]int) bool {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for _, num := range counts {
		if num > 1 {
			return true
		}
	}
	return false
}
