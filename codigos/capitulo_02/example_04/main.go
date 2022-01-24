package main

import "fmt"

func main() {
	var myvar float32
	fmt.Print("Introduce un numero: ")
	fmt.Scanf("%f", &myvar)
	fmt.Println(myvar)
	fmt.Println("Mi variable es:", myvar)
	fmt.Printf("Mi variable es: %f\n", myvar)
}
