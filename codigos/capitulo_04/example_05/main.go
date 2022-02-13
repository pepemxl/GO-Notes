package main

import "fmt"

func function_clousure() {
	a := 5
	lambda_style := func() {
		lambda_style2 := func() {
			fmt.Println("Funcion lambda_style2", a)
		}
		lambda_style2()
	}
	lambda_style()
}

func main() {
	function_clousure()
}
