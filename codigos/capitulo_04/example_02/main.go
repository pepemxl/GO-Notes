package main

import "fmt"

func compute(a int, b int, c int) (res1 int, res2 float32, res3 float32) {
	res1 = a + b + c
	res2 = float32(a * b * c)
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	return res1, res2, res3
}

func main() {
	compute(1, 2, 3)
}
