package main

import (
	list "data_structure/Data/Static_list"
	"fmt"
)

func main() {
	l := list.New().Init(666)
	l.Append(555)
	l.Append(66666)
	l.Append(55)
	l.Append(22)
	l.InsertAfter(333, 2)
	l.InsertBefore(222, 2)
	l.FirstInsert(900)
	fmt.Println(l)

}
