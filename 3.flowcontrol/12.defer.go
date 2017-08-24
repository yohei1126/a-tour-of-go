// https://tour.golang.org/flowcontrol/12
// Defer
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
