package main

import "fmt"

// func add(x, y int) (int, int) {
// 	//fmt.Printf("Sum is : %d", x+y)
// 	return x + y, x - y
// }

// func add(x, y int) (z1 int, z2 int) { //z1 and z2 are labels
// 	z1 = x + y
// 	z2 = x - y
// 	return
// }

//Using Defer keyword

func add(x, y int) (z1 int, z2 int) {
	defer fmt.Println("Hello")
	z1 = x + y
	z2 = x - y
	fmt.Println("Before return") //this prints first
	return
}

func main() {
	// ans1, ans2 := add(3, 53)
	// fmt.Println(ans1, "\n", ans2)
	x := add
	ans1, ans2 := x(3, 53)
	fmt.Println(ans1, ans2)
}
