package main

import "fmt"

func main() {
	var i int
	fmt.Println("for i = 0; i < 10; i++ {")
	fmt.Println("	fmt.Printf(\"%d \", i+1)")
	fmt.Println("}")
	for i = 0; i < 10; i++ {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println("\n\nfor ii := 0; ii < 10; ii++ {")
	fmt.Println("	fmt.Printf(\"%d \", ii+1)")
	fmt.Println("}")
	for ii := 0; ii < 10; ii++ {
		fmt.Printf("%d ", ii+1)
	}
	fmt.Println("\n\nk := 0")
	fmt.Println("for k < 10 {")
	fmt.Println("	fmt.Printf(\"%d \", k+1)")
	fmt.Println("}")
	k := 0
	for k < 10 {
		fmt.Printf("%d ", k+1)
		k++
	}
}
