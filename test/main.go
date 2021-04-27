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

	value := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(value[:0])
}
