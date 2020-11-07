package main

import (
	"fmt"
)

func main() {
	// x := 5
	// y := 6.5
	x := "Mit"
	y := "mit"
	//val := float64(x) == y //Other operators are same as python
	val := x == y
	fmt.Printf("%t\n", val)
	fmt.Printf("%t\n", x != y)
}
