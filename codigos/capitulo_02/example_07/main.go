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
}
