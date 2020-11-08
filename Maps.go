package main

import "fmt"

func main() {
	//map does not keep track of order
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   6,
		"orange": 9,
	}
	// fmt.Println(mp["apple"])
	// mp["apple"] = 9
	// fmt.Println(mp["apple"])
	// fmt.Println(mp)
	// //mp := make(map[string]int)  // another method to make a map
	// mp["Mit"] = 777
	// delete(mp, "apple")

	val, ok := mp["apple"]
	val1, ok := mp["Mit"]
	fmt.Println(val, ok)
	fmt.Println(val1, ok)
	fmt.Println(mp)
	fmt.Println(len(mp))
}
