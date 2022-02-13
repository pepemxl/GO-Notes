package main

import (
	"fmt"
)

func main() {
	fmt.Println("Example 02")
	var arr [5]string
	for i := 0; i < 5; i++ {
		arr[i] = fmt.Sprintf("Mi cadena %d", i+1)
	}
	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}
}
