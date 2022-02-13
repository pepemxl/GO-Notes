package main

import "fmt"

func main() {
	var option int = 3
	switch option {
	case 1:
		fmt.Println("Option 1")
	case 2:
		fmt.Println("Option 2")
	case 3:
		fmt.Println("Option 3")
	default:
		fmt.Println("Option por defecto")
	}
	x := 8
	switch y := x % 2; y {
	case 0:
		x -= 1
	case 1:
		x += 1
	}
	fmt.Println(x)
}
