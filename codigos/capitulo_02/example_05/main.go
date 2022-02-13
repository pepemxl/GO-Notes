package main

import "fmt"

func main() {
	var a, b int = 2, 10
	fmt.Println(" a > b ", a > b)
	fmt.Println(" a < b ", a < b)
	fmt.Println(" a >= b ", a >= b)
	fmt.Println(" a <= b ", a <= b)
	fmt.Println(" a != b ", a != b)
	fmt.Println(" a == b ", a == b)
	fmt.Println(" true && true ", true && true)
	fmt.Println(" true && false ", true || false)
	fmt.Println(" !true ", !true)
	fmt.Println(" 1 & 2 ", 1&2)
	fmt.Println(" 1 & 3 ", 1&3)
	fmt.Println(" 2 >> 1 ", 2>>1)
	fmt.Println(" 2 << 1 ", 2<<1)
}
