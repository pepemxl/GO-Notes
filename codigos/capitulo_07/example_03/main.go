package main

import "fmt"

func main() {
	num1 := 1
	num2 := 3.14159
	var str_num1 string
	var str_num2 string
	var str_num3 string
	str_num1 = fmt.Sprintf("%d", num1)
	str_num2 = fmt.Sprintf("%d", num2)
	str_num3 = fmt.Sprintf("%f", num2)
	fmt.Println(str_num1)
	fmt.Println(str_num2)
	fmt.Println(str_num3)
}
