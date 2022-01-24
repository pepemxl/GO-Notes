package main

import "fmt"

func main() {
	var a, b int = 1, 2
	if a > b {
		fmt.Printf("%d > %d", a, b)
	} else {
		fmt.Printf("%d <= %d", a, b)
	}
}
