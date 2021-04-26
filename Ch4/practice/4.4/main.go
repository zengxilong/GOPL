package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(rotate(a, 4))
}
func rotate(slice []int, position int) []int {
	res := slice[position:]
	for i := 0; i < position; i++ {
		res = append(res, slice[i])
	}
	return res
}
