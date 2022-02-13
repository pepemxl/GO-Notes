package main

import "fmt"

func hello_world_str() string {
	return "Hello World"
}

func hello_world_print() {
	fmt.Println("Hello World 2")
}

func hello_world_print_parameter(_val int) {
	fmt.Println("Hello World", _val)
}

func hello_world_print_multi_parameter(_val1, _val2, _val3 int) {
	fmt.Println("Hello World", _val1, _val2, _val3)
}

func hello_world_print_multi_parameter2(_val1 int, _val2 int, _val3 int) {
	fmt.Println("Hello World", _val1, _val2, _val3)
}

func main() {
	fmt.Println(hello_world_str())
	hello_world_print()
	hello_world_print_parameter(3)
	hello_world_print_multi_parameter(1, 2, 3)
	hello_world_print_multi_parameter2(1, 2, 3)
}
