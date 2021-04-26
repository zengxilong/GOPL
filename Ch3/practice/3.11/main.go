package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(comma("-312323.145142"))
}

func commaBase(str string) string {
	var buf bytes.Buffer
	for i := 0; i < len(str); i++ {
		if (len(str)-i)%3 == 0 && i != 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(str[i])
	}
	return buf.String()
}

func comma(s string) string {
	if s == "" {
		return s
	}
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	var afterDot, beforeDot string
	for i, v := range s {
		if v == '.' {
			afterDot = s[i:]
			beforeDot = s[:i]
		}
	}

	buf.WriteString(commaBase(beforeDot))
	buf.WriteString(afterDot)
	return buf.String()
}
