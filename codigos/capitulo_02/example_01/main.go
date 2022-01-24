package main

import "fmt"

func main() {
	// Comments lines using // or  for multiline comments /* */
	var myvar1 string      // declaración de nuestra variable
	myvar1 = "Hello World" // asignación de valor
	fmt.Println(myvar1)
	myvar1 = "Hello World 2" // reasignación
	fmt.Println(myvar1)
	myvar2 := "Hello World 3" // declaracion utilizando :=
	fmt.Println(myvar2)
	myvar2 = "Hello World 4"
	fmt.Println(myvar2)
	var myvar3 = "Hello World 5" // declaracion y asignación en la misma linea no necesita tipo
	fmt.Println(myvar3)
	myvar3 = "Hello World 6"
	fmt.Println(myvar3)
	var myvar4 = 7
	fmt.Println("Hello World", myvar4)
	var ( // declaracion de variable en grupo de distintos tipos
		name  string
		email string
		age   int
	)
	name = "pepe"
	email = "pepemxl@gmail.com"
	age = 37
	fmt.Println(name, email, age)
	const pi = 3.14159 // variables constantes
	fmt.Println("PI: ", pi)
}
