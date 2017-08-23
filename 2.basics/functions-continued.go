// https://tour.golang.org/basics/5
// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
package main

import "fmt"

// omit the type from all but the last.
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
