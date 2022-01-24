package main

import "fmt"

func main() {
	var i int
	fmt.Println("for i = 0; i < 10; i++ {")
	fmt.Println("	if i == 5 {")
	fmt.Println("		break")
	fmt.Println("	}")
	fmt.Println("	fmt.Printf(\"%d \", i+1)")
	fmt.Println("}")
	for i = 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("%d ", i+1)
	}
	fmt.Println("")
	fmt.Println("for i = 0; i < 10; i++ {")
	fmt.Println("	if i == 5 {")
	fmt.Println("		continue")
	fmt.Println("	}")
	fmt.Println("	fmt.Printf(\"%d \", i+1)")
	fmt.Println("}")
	for i = 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%d ", i+1)
	}
}
