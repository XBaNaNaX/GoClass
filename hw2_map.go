package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var m map[string]int
	m = make(map[string]int)

	sf := strings.Fields(s)
	for i := 0; i < len(sf); i++ {
		if m[sf[i]] == 0 {
			m[sf[i]] = 1
		} else {
			m[sf[i]] = m[sf[i]] + 1
		}
	}
	return m
}
func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck"

	w := WordCount(s)
	fmt.Println(w)
}
