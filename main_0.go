// package main

// import "fmt"

// func main() {
// var number string= "hel"
// var num int = 250
// var number int
// var bl bool
// bl = true
// fmt.Println(number, bl)
//fmt.Printf("%T", number)
// fmt.Printf("%T", 10)
// fmt.Printf("Hello %T %v", 10, 10)
//}

package main

import "fmt"

func main() {
	num := 3240000
	num1 := 2.335232135245426462
	fmt.Printf("Binary: %b", num, "\n")
	fmt.Printf("Octal: %o", num, "\n")
	fmt.Printf("Decimal: %d", num, "\n")
	fmt.Printf("Hexadecimal: %x", num, "\n")
	fmt.Printf("Hexadecimal: %X", num, "\n")
	fmt.Printf("Scientific notation: %e", num1, "\n")
	fmt.Printf("Limits to certain digit:%f", num1, "\n")
	fmt.Printf("large decimal: %g", num1, "\n")
	fmt.Printf("Name: %s", "Mit", "\n")
	fmt.Printf("Name with double quotation: %q", "Mit", "\n")
	fmt.Printf("Name:%9", "Mit", "\n") //padding in python string_name.ljust(length_of_string_as_final_output," ")
	fmt.Printf("Number: %07d is awesome", 45, "\n")
	var str1 string = fmt.Sprintf("Number: \n %07d is awesome", 45)
	fmt.Println(str1)
	var str0 string = fmt.Sprintf("Number: \t %07d is awesome", 45)
	fmt.Println(str0)
}
