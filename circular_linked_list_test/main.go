package main

import (
	list "data_structure/Data/circular_linked_list"
	"fmt"
)

func main() {
	l := list.New().Init(0)
	l.BackInsert(1)
	l.BackInsert(2)
	l.BackInsert(3)
	print(l)
	l.InsertNode(4, 0)
	print(l)
	l.InsertNode(4, 1)
	print(l)
	l.BackInsert(99)
	print(l)
	l.FristInsert(1234)
	print(l)
}
func print(l *list.List) {
	node := l.HeadNode()
	fnode := l.TailNode()
	fmt.Print(node.Get(), "-->")
	for node != fnode {
		node = node.Next()
		fmt.Print(node.Get(), "-->")
	}
	fmt.Println()

}
