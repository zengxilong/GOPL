package main

import (
	"bytes"
	"fmt"
)

func comma(str string) string {
	var buf bytes.Buffer
	for i := 0; i < len(str); i++ {
		if (len(str)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(str[i])
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("1234567"))
}
