package main

import (
	"data_structure/Data/myQueue"
	"fmt"
)

func main() {
	q := myQueue.NewQueue()
	q.Push(10)
	q.Push(25)
	q.Push(22)
	fmt.Println(q.Len())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())

}
