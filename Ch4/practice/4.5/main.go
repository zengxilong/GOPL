package main

import "fmt"

func removeDup(slice []string) []string {
	for i := 0; i < len(slice)-1; {
		if slice[i] == slice[i+1] {
			copy(slice[i:], slice[i+1:])
			slice = slice[:len(slice)-1]
		} else {
			i++
		}
	}
	return slice
}

func main() {
	fmt.Println(removeDup([]string{"hello", "hello", "hello", "nihao"}))
}
