package main

import "strings"

func ReverseSentence(s string) string {
	if s == "" {
		return ""
	} else {
		words := strings.Split(s, " ")
		for i := 0; i < len(words)/2; i++ {
			words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
		}
		return strings.Join(words, " ")
	}
}

func main() {
	str := "Разворот слов в предложении"

	println(ReverseSentence(str))
}
