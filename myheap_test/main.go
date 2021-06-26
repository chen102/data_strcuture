package main

import "data_structure/Data/myheap"
import "fmt"

//import "container/heap"

func main() {
	arr := []int{4, 6, 8, 5, 9}
	h := myheap.New(arr)
	h.Push(66)
	fmt.Println(h.Peek())
	//heap.Init()

}
