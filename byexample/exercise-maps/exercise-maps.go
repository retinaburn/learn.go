package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		count, ok := m[word]
		if !ok {
			m[word] = 1
		} else {
			m[word] = count + 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
