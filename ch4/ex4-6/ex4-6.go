package main

import (
	"fmt"
	"unicode"
)

func main() {
	d := []byte("abc     a aaa     ccc  ddd d")
	e := RemoveSpace(d)
	fmt.Println(string(e))
}

func RemoveSpace(s []byte) []byte {
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			// 结合remove函数
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
