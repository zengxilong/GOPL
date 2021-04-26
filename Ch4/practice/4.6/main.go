package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func replace(b []byte) []byte {
	for i := 0; i < len(b); {
		front, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(front) {
			after, _ := utf8.DecodeRune(b[i+size:])
			if unicode.IsSpace(after) {
				copy(b[i:], b[i+size:])
				b = b[:len(b)-size]
			}
		}
		i += size
	}
	return b
}

func main() {
	b := []byte("你  界的")

	fmt.Printf("%s\n", replace(b))
}
