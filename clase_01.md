# GO


El lenguaje de programación GO es muy similar a los lenguajes de la familia "C". Una diferencia significativa es que no necesitamos usar `;`.


Para la declaracion de variables solo necesitamos usar la siguiente sintaxis:

```GO
var myvar data_type
```

por ejemplo:

```GO
var i int
```
 o varias variables en la misma linea:

```GO
var i, j int
```

podemos hacer la asignación en una misma linea, aunuque la sintaxis queda un tanto extraña

```GO
var i, j int = 1, 2
```
aunque go puede interpretar el tipo si realizamos la asignación al momento de la declaración:

```GO
var i, j = 1, 2
```

## Tipos de datos en Go

- int - integer
- float32 - a single-precision floating point value.
- float64 - a double-precision floating point value.
- string - a text value.
- bool - Boolean true or false.


## Operadores incremento `++` y decremento `--`

Los operadores de incremento/decremento funcionan casi igual que en C/C++ sin embargo tenemos pequeños casos en los cuales no funciona, como por ejemplo

```go
fmt.Printf("a = %d\n", a++)
```
falla durante la compilación.

## `Printf` y `Scanf`

Las funciones `Printf` y `Scanf` funcionan igual que sus contrapartes en lenguajes como C/C++.

## Operadores logicos

Los operadores logicos `&&`, `||` y `!` funcionan al igual que lo haciamos en C/C++.

## Operador de control condicional `if`

La sintaxis es 

```go
if conditional {
    // do something in true case
}else{
    // do something in false case
}
```

una cosa a resaltar es que el copilador y/o los snippets terminan quitando los parentesis, sin embargo aqui las llaves son requeridas!!!

## Operador de control `switch`

El operador `switch` funciona igual que en otros lenguajes:

```go
package main

import "fmt"

func main() {
	var option int = 3
	switch option {
	case 1:
		fmt.Println("Option 1")
	case 2:
		fmt.Println("Option 2")
	case 3:
		fmt.Println("Option 3")
	default:
		fmt.Println("Option por defecto")
	}
}
```



















