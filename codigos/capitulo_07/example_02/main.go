package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "1"
	str2 := "2"
	var err error
	var val1 int64
	val1, err = strconv.ParseInt(str1, 10, 32)
	if err == nil {
		fmt.Println(val1)
	} else {
		fmt.Println(err)
	}
	var val2 float64
	val2, err = strconv.ParseFloat(str2, 64)
	if err == nil {
		fmt.Println(val2)
	} else {
		fmt.Println(err)
	}

}
