package main

import (
	list "data_structure/Data/Double_Linked_Lists"
	"fmt"
)

func main() {
	l := list.New().Init(0)
	root := l.Root()
	node1 := l.InsertAfter(3, root)
	node2 := l.InsertBefore(2, node1)
	fmt.Println(node1.GetValue())
	fmt.Println(node2.GetValue())
	//n := root.Next()
	//for n != nil {
	//fmt.Println(n.GetValue())
	//n = n.Next()
	//}

	fmt.Println(l.Front().GetValue())
	l = l.Remove(node2)
	fmt.Println(l.Front().GetValue())
}
