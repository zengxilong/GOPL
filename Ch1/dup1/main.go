package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	content := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		content[input.Text()]++
	}

	for line, num := range content {
		if num > 1 {
			fmt.Printf("%s\t%d\n", line, num)
		}
	}
}
