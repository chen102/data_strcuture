package main

import (
	"data_structure/Data/Stack"
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	//后缀表达式计算结果
	exp := "931-3*+102/+" //本质上是[]byte{}
	s := Stack.NewStack1()
	for _, v := range exp { //使用for range遍历字符串 返回rurn类型(int32类型)
		if unicode.IsDigit(v) {
			if runeToint(v) == 0 { //是数字10入栈而不是1和0入栈
				s.Pop()
				s.Push(10)
				continue
			}
			s.Push(runeToint(v))

		} else {
			res := cal(s.Pop().(int), s.Pop().(int), v)
			s.Push(res)
		}
	}
	fmt.Println(s.Peek())

	//中缀表达式转后缀表达式
	exp1 := "9+(3-1)*3+10/2"
	s1 := Stack.NewStack1()
	for _, v := range exp1 {
		if unicode.IsDigit(v) {
			fmt.Printf("%c", v)
		} else {
			if s1.Len() == 0 {
				s1.Push(string(v))
				continue
			}
			for s1.Len() > 0 && priority(string(v), s1.Peek().(string)) {
				if string(v) == ")" && s1.Peek().(string) == "(" {
					s1.Pop()
					break
				}
				fmt.Printf(s1.Pop().(string))

			}
			if string(v) == ")" {
				continue
			}
			s1.Push(string(v))

		}
	}
	for s1.Len() != 0 {
		fmt.Printf(s1.Pop().(string))
	}

}
func cal(a, b int, exp rune) (res int) {
	switch exp {
	case '+':
		res = b + a
	case '-':
		res = b - a
	case '*':
		res = b * a
	case '/':
		res = b / a
	}
	return
}
func runeToint(s rune) (res int) {
	res, err := strconv.Atoi(string(s))
	if err != nil {
		panic(err)
	}
	return

}
func priority(a, b string) (res bool) {
	switch a {
	case "+", "-":
		if b == "+" || b == "-" || b == "*" || b == "/" {
			res = true
		}
	case "/", "*":
		if b == "*" || b == "/" {
			res = true
		}
	case ")":
		res = true

	}
	return
}
