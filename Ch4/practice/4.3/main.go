package main

import "fmt"

func reverse(ints *[10]int) {
	for i, j := 0, len(*ints)-1; i < j; i, j = i+1, j-1 {

		ints[i], ints[j] = ints[j], ints[i]

	}
}

func main() {
	var ints = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	reverse(&ints)
	fmt.Println(ints)
}
