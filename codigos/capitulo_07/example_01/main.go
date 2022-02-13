package main

import "fmt"

func main() {
	var myStr string
	myStr = "Hello World"
	fmt.Println(myStr)
	var str1 string
	var str2 string
	str1 = "Hello"
	str2 = "World"
	str3 := str1 + str2
	fmt.Println(str3)
	str4 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str4)
}
