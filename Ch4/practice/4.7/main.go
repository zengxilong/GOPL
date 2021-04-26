package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseUTF8(b []byte) {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
	reverse(b)
}

func main() {
	b := []byte("你好世界")
	reverseUTF8(b)
	fmt.Printf("%s\n", b)
}
