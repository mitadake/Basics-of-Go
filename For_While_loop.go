package main

import (
	"fmt"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }
	// for x := 0; x < 1000; x++ {
	// 	fmt.Println(x)
	// } //Output will be same as above
	//for{} is a valid syntax and acts same as while loop
	//syntax for ewhile loop is while{} which is same as any other language
	for x := 0; x < 1000; x++ {
		if x != 0 && x%3 == 0 && x%7 == 0 && x%9 == 0 {
			fmt.Println(x)
			//continue and break keyword valid in GO
		}
	}
}
