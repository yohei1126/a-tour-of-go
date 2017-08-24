// https://tour.golang.org/flowcontrol/2
// For continued
package main

import "fmt"

func main() {
	sum := 1
	// The init and post statement are optional.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
