package main

import "fmt"

func main() {
	var x int = 5
	var ptr_x *int = &x
	var y *int
	y = new(int)
	*y = 6
	fmt.Println("x=", x)
	fmt.Println("&x=", &x)
	fmt.Println("ptr_x=", ptr_x)
	fmt.Println("y=", y)
	fmt.Println("*y=", *y)
}
