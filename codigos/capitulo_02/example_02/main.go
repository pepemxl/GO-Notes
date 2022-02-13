package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	var d float32
	a = 1
	b = 2
	c := a + b
	d = 3.14159
	e := float32(a) + d
	fmt.Printf("%d + %d = %d\n", a, b, c)
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("float32(%d + %d) = %.2f\n", a, b, float32(a+b))
	fmt.Printf("sin(%d + %d) = %.2f\n", a, b, math.Sin(float64(a+b)))
	fmt.Println(e)
}
