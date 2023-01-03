package main

import "fmt"

func main() {
	x := 15
	a := &x // & ->memory address *-> read whats at that memory address  

	fmt.Println(a)
	fmt.Println(*a)
	*a = 5   // changes 

	fmt.Println(x)
	*a = *a**a // squaring the value at x 
	fmt.Println(x)
	fmt.Println(*a)
}