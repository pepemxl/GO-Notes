package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	var a, b = 1, 2
	a, b = swap(a, b)
	fmt.Println(a, b)
	c, d := swap(1, 2)
	fmt.Println(c, d)
}
