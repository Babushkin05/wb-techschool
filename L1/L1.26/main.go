package main

import (
	"fmt"
	"unicode"
)

func IsUnique(s string) bool {
	m := map[rune]struct{}{}

	runes := []rune(s)
	for _, r := range runes {
		if _, ok := m[unicode.ToLower(r)]; ok {
			return false
		} else {
			m[r] = struct{}{}
		}
	}
	return true
}

func main() {
	for _, s := range []string{"abcd", "abCdefAf", "aabcd"} {
		fmt.Printf("%s: %v\n", s, IsUnique(s))
	}
}
