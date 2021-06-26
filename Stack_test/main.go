package main

import (
	"data_structure/Data/Stack"
	"fmt"
	"math"
)

func main() {
	//s := Stack.NewStack(10)
	//s1 := Monotonous([]int{4, 2, 22, 2, 9, 5}, s)
	//for !s1.NullStack() {
	//fmt.Printf("%d,", s1.Pop())
	//}

	//s2 := Stack.NewStack1()
	//s2.Push(22)
	//s2.Push(21)
	//fmt.Println(s2.Peek().(int))

	//s3 := Stack.NewStack(10)
	//fmt.Println(Monotonous1([]int{4, 3, 7, 1}, s3))
	//s := Monotonous1([]int{4, 3, 7, 1}, s3)
	//fmt.Println()
	//fmt.Println(s)

	//s4 := Stack.NewStack(10)
	//s5 := Monotonous2([]int{2, 1, 5, 6, 2, 3}, s4)
	//fmt.Println(s5)

	//DoubleStack := Stack.NewDoubleStack(20)
	//DoubleStack, err := DoubleStack.Push(1, 1)
	//if err != nil {
	//fmt.Println(err)
	//}
	//DoubleStack, err = DoubleStack.Push(22, 2)
	//if err != nil {
	//fmt.Println(err)
	//}

	//fmt.Println(DoubleStack.Ele)
	//fmt.Println(DoubleStack.Pop(1))
	//fmt.Println(DoubleStack.Pop(1))
	//fmt.Println(DoubleStack.Pop(1))
	//fmt.Println(DoubleStack.Pop(2))
	//fmt.Println(DoubleStack.Pop(2))
}

//单调栈
//4进栈，栈为空，直接进栈,栈内元素为4
//2进栈，栈顶为4，比2大，入栈，栈内元素为4,2
//22进栈,栈顶为2,比22小，出栈，栈顶元素为4,比22小，出栈，栈顶为空，将22进栈,栈内元素为22
//2进栈,栈顶为22,比2大，进栈，栈内元素为22,2
//.......
func Monotonous(a []int, s *Stack.Stack) *Stack.Stack {
	for _, v := range a {
		if s.NullStack() || s.Peek().(int) >= v {
			s.Push(v)
			fmt.Println(v, "已入栈")
		} else {
			for !s.NullStack() && s.Peek().(int) <= v {
				fmt.Println(s.Pop(), "已出栈")
			}
			s.Push(v)
			fmt.Println(v, "已入栈")

		}
	}
	return s
}

//1.视野总和
//描叙：有n个人站队，所有的人全部向右看，个子高的可以看到个子低的发型，给出每个人的身高，问所有人能看到其他人发现总和是多少。
//实质上是找当前数字向右查找的第一个大于他的数字之间有多少个数字，然后将每个结果累计
//使用单调栈来解决这个问题
func Monotonous1(a []int, s *Stack.Stack) int {
	sum := 0
	a = append(a, math.MaxInt64) //需要一个无线高的人档住栈中的人,不然栈中元素最后无法完全出栈
	for k, v := range a {
		if s.NullStack() || a[s.Peek().(int)] > v {
			s.Push(k)
		} else {
			for !s.NullStack() && a[s.Peek().(int)] <= v {
				top := s.Peek().(int)
				s.Pop()
				sum += (k - top - 1) //第一个比他大的数字的位置减去前面数字的位置
				fmt.Println("sum:", sum)
			}
			s.Push(k)

		}
	}
	return sum

}

//算法流程
//0进栈，栈为空，直接进栈，栈内元素为0
//1进栈,栈顶为0，a[0]=4比a[1]=3大，入栈，栈内元素0,1
//2进栈,栈顶为1，a[1]=3比a[2]=7小，出栈，更新sum:sum+=2-1-1,栈内元素0
//2进栈,栈顶为0，a[0]=4比a[2]=7小,出栈，更新sum:sum+=2-0-1，空栈了退出循环2入栈，栈内元素2
//3进栈，栈顶为2,a[2]=7比a[3]=1大,入栈,栈内元素2,3
//4进栈,栈顶为3，a[3]=1比a[4]=MAX小，出栈,更新sum:sum+=4-3-1，栈内元素2
//4进栈,栈顶为2，a[2]=7比a[4]=MAX小，出栈,更新sum:sum+=4-2-1，栈内元素2

//柱状图中的最大矩形
//给定n个非负数整数，用来表示柱状图中各个柱子的高度,每个柱子彼此相邻，且宽度为1，求在该柱状图能够勾勒出来的矩形的最大面积
//这里的宽度实质上是上一题中的与最大的数中间有多少数字是一样的
func Monotonous2(a []int, s *Stack.Stack) int {
	a = append(a, -1) //同样需要所有元素出栈
	ret, top := 0, 0
	for k, v := range a {
		if s.NullStack() || a[s.Peek().(int)] <= v {
			s.Push(k)
		} else {
			for !s.NullStack() && a[s.Peek().(int)] > v { //牢记栈中数据永远是有序的
				top = s.Peek().(int)
				s.Pop()
				tmp := (k - top) * a[top]
				if tmp > ret {
					ret = tmp
				}
			}
			s.Push(top)
			a[top] = v

		}
	}
	return ret
}

//算法流程
//栈为空0直接进栈，栈内元素0
//1进栈，栈顶为0,a[0]=2>a[1]=1,出栈，计算面积tmp=(1-0)*2=2大于ret，更新ret,栈为空退出循环，0重新进栈但a[0]=1,栈内元素0
//2进栈，栈顶为0,a[0]=1<a[2]=5,入栈，栈内元素0，2
//3进栈，栈顶为2,a[2]=5<a[3]=6,入栈，栈内元素0，2，3
//4进栈，栈顶为3,a[3]=6>a[4]=2,出栈,计算面积tmp=(4-3)*6=6,大于ret，更新ret,栈顶2，a[2]=5>a[4]=2,出栈，计算面积tmp=(4-2)*5=10大于ret，更新ret,栈顶为1，a[1]=1<a[4]=2，退出循环，2重新进栈但a[2]=2,栈内元素1，2
//5进栈,栈顶为2，a[2]=2<a[5]=3，入栈,栈内元素1,2,5
//6进栈，栈顶为5，a[5]=3>a[6]=-1，出栈，计算面积tmp=(6-5)*5=5小于ret,栈顶为2，a[2]=2>a[6]=-1，出栈，计算面积tmp=(6-2)*2=8小于ret，栈顶为1，a[1]=1>a[6]=-1，出栈，计算面积tmp=(6-1)*5=10等于ret，栈顶为空遍历结束返回最大面积10
