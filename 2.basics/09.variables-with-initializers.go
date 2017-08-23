// https://tour.golang.org/basics/9
// Variables with initializers
package main

import "fmt"

var i, j = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
