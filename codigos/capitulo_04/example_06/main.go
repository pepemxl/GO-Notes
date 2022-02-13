package main

import "fmt"

func main() {
	fmt.Println(fibo(5))
}

func fibo(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}
