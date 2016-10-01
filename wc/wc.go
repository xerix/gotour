package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	char := strings.Fields(s)
	for i := range char {
		c := char[i]
		_, ok := m[c]
		if ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
