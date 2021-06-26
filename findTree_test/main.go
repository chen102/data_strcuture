package main

import (
	tree "data_structure/Data/findTree"
	"fmt"
)

func main() {
	root := tree.New(6)
	root.Insert(2)
	root.Insert(7)
	root.Insert(1)
	fmt.Println(root.Left.E)
	fmt.Println(root.Right.E)
	fmt.Println(root.Left.Left.E)
	fmt.Println(root.FindMin().E)
	fmt.Println(root.FindMax().E)
	root.Find(9)
}
