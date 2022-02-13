package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "manzana;pera;;naranja;"
	mylista := strings.Split(data, ";")
	for idx, fruta := range mylista {
		fmt.Println(idx, fruta, len(fruta))
	}
	fmt.Println(data[0:len(data)])
	fmt.Println(data[0:7])
	fmt.Println(data[:7])
	fmt.Println(data[7:])
	fmt.Println(data[len(data)-5:])
	fmt.Println(strings.ToLower(data[0:7]))
	fmt.Println(strings.ToUpper(data[0:7]))
}
