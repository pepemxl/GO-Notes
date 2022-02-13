package main

import "fmt"

func compute(a int, b int, c int) (res1 int, res2 float32, res3 float32) {
	res1 = a + b + c
	res2 = float32(a * b * c)
	// res3 = float32(a / res2) //error
	res3 = float32(a) / res2
	return res1, res2, res3
}

func main() {
	var res1 int
	var res2, res3 float32
	res1, res2, res3 = compute(1, 2, 3)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
}
