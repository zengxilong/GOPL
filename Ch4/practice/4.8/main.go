package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	letter  string = "letter"
	number         = "number"
	symbol         = "symbol"
	graphic        = "graphic"
	space          = "space"
)

func main() {
	counts := make(map[string]int)

	in := bufio.NewReader(os.Stdin)

	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		switch {
		case unicode.IsLetter(r):
			counts[letter]++
		case unicode.IsNumber(r):
			counts[number]++
		case unicode.IsSymbol(r):
			counts[symbol]++
		case unicode.IsGraphic(r):
			counts[graphic]++
		case unicode.IsSpace(r):
			counts[space]++
		}
	}
	fmt.Printf("class\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}

}
