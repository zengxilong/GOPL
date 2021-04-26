package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("hello", "llohe"))
}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	charMap := make(map[rune]int, len(str1))

	for _, v := range str1 {
		charMap[v]++
	}

	for _, v := range str2 {
		if charMap[v] == 0 {
			return false
		}
		charMap[v]--
	}
	return true
}
