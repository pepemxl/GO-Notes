package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func add(list *Node, data int) *Node {
	if list == nil {
		list := new(Node)
		list.value = data
		list.next = nil
		return list
	} else {
		temp := new(Node)
		temp.value = data
		temp.next = list
		return temp
	}
}

func display(list *Node) {
	var temp *Node
	temp = list
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func main() {
	fmt.Println("Hello")
	var lista *Node
	lista = nil
	for i := 0; i < 10; i++ {
		lista = add(lista, i+1)
	}
	display(lista)
}
