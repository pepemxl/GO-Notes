package main

import "fmt"

func main() {
	var myArray [10]int
	for i := 0; i < 10; i++ {
		myArray[i] = 10*i + 1
	}
	for i := 0; i < 10; i++ {
		fmt.Println(myArray[i])
	}
	var myArray2 [3][3]float32
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			myArray2[i][j] = float32(i)*3 + float32(j)
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(myArray2[i][j], " ")
		}
		fmt.Println()
	}
}
