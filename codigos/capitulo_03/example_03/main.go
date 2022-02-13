package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// define slice
	fmt.Println("define slices")
	var numbers []int
	numbers = make([]int, 100)
	for i := 0; i < 100; i++ {
		numbers[i] = rand.Intn(100)
	}
	for i := 0; i < 100; i++ {
		fmt.Print(numbers[i])
	}
	fmt.Println()
	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		matrix[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			matrix[i][j] = i*3 + j
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

}
