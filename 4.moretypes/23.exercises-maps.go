// https://tour.golang.org/moretypes/23
// Exercise: Maps
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, f := range strings.Fields(s) {
		m[f]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
