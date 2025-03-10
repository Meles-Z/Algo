package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
func main() {
	fmt.Println(reverseString("one two three"))
}
