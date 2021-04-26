package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCount := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		wordCount[input.Text()]++
	}

	fmt.Printf("count\tword\n")
	for c, n := range wordCount {
		fmt.Printf("%d\t%s\n", n, c)
	}
}
