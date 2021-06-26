package main

import (
	ChildShow "data_structure/Data/child_show"
	"fmt"
)

func main() {
	l := ChildShow.New("A")
	l.Add("B", 0)
	l.Add("C", 1)
	l.Add("D", 1)
	l.Add("E", 0)
	fmt.Println(l.Node[0].Firstchild.Next.Child)
}
