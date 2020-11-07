package main

import (
	"fmt"
)

func main() {
	n := 13
	if n >= 18 {
		fmt.Println("Good to go!")
	} else if n >= 14 {
		fmt.Printf("U Can go ahead with assistance\n")
	} else {
		fmt.Printf("U Can't go ahead\n")
		fmt.Printf("Extra %d years required!! ", 18-n)
	}
}
