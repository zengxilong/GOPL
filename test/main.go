package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, 世界"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	fmt.Println(len(s))
	r, size := utf8.DecodeRuneInString("世界")
	fmt.Printf("%d\t%c\n", size, r)
	fmt.Println(utf8.RuneCountInString(s))
}
