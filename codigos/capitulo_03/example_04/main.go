package main

import "fmt"

func main() {
	var myjson = make(map[string]int)
	myjson["uno"] = 1
	myjson["dos"] = 2
	myjson["tres"] = 3
	myjson["cuatro"] = 4
	myjson["cinco"] = 5
	fmt.Println(myjson["uno"])
	fmt.Println(myjson["dos"])
	fmt.Println(myjson["tres"])
	fmt.Println(myjson["cuatro"])
	fmt.Println(myjson["cinco"])
	var myjson2 = make(map[int]string)
	myjson2[1] = "uno"
	myjson2[2] = "dos"
	myjson2[3] = "tres"
	myjson2[4] = "cuatro"
	myjson2[5] = "cinco"
	for i := 1; i <= 5; i++ {
		fmt.Println(myjson2[i])
	}
	fmt.Println(myjson)
	fmt.Println(myjson2)
}
