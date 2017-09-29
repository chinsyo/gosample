package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		word := words[i]
		count, exist := result[word]
		if exist {
			result[word] = count + 1
		} else {
			result[word] = 1
		}
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
