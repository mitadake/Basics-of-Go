package main

import (
	"fmt"
)

func main() {
	num1 := 8.0 //var num1 float64 := 8
	num2 := 10.0
	answer := num1 / num2
	ans1 := num1 + num2
	ans2 := num1 - num2
	ans3 := num1 * num2
	ans4 := int(num1) / int(num2)
	num3 := 8
	num4 := 10
	ans5 := num3 / num4
	ans6 := num3 % num4
	fmt.Printf("%g ", answer)
	fmt.Printf("%g ", ans1)
	fmt.Printf("%g ", ans2)
	fmt.Printf("%g ", ans3)
	fmt.Printf("%d ", ans4)
	fmt.Printf("%d ", ans5)
	fmt.Printf("%d ", ans6)
}
