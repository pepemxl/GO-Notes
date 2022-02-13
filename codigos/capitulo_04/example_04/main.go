package main

import "fmt"

func unlimited(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func main() {
	result := unlimited(1, 2, 3)
	fmt.Println(result)
}
