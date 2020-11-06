package main

import (
	"fmt"
)

func main() {
	// var arr [5]string
	// fmt.Println(arr)
	// var arr [5]int
	// fmt.Println(arr)
	// var arr [5]int
	// arr[0] = 10
	// arr[4] = 9
	//fmt.Println(arr[0])
	// fmt.Println(arr[0])
	/*Another method*/
	// arr := [3]int{2, 41, 24}
	// fmt.Println(arr)
	// fmt.Println(len(arr))
	// arr := [3]int{2, 41, 24}
	// sum := 0
	// for i := 0; i < len(arr); i++ {
	// 	sum += arr[i]
	// }
	// fmt.Println(sum)
	arr2D := [2][3]int{{1, 2, 4}, {3, 4, 6}}
	arr2D[0][1] = 0
	fmt.Println(arr2D)
}
