// https://tour.golang.org/basics/14
// Type inference

package main

import "fmt"

func main() {
	//v := 42 // int
	v := "42" // string
	fmt.Printf("v is of type %T\n", v)
}
