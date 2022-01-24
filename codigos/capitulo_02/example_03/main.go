package main

import "fmt"

func main() {
	a := 1
	fmt.Printf("a = %d\n", a)
	a++
	fmt.Printf("a = %d\n", a)
	a--
	fmt.Printf("a = %d\n", a)
	//fmt.Printf("a = %d\n", a++)// NO compila
	//fmt.Printf("a = %d\n", (a--)) // NO compila
	fmt.Printf("a = %d\n", a-1) // SI compila
}
