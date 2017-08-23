// https://tour.golang.org/basics/4
// A function can take zero or more arguments.
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
