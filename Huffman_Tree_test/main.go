package main

import (
	"data_structure/Data/myQueue"
	"data_structure/Data/tree"
	"fmt"
	"math/rand"
	//"time"
)

//哈夫曼树(最优树):带权路径长度(WPL)最短的树(度相同的情况下)
//贪心算法:构造哈夫曼树时首先选择权值小的叶子节点

var c []int

const MAXTREE = 5

type binaryforest []*tree.BinaryTree

func main() {
	forest := make([]*tree.BinaryTree, 0, MAXTREE)
	for i := 0; i <= MAXTREE; i++ { //构造节点全是根
		binarytree := tree.NewBinaryTree().InitBinaryTree(rand.Intn(10))
		forest = append(forest, binarytree)
	}
	tree := BuildHuffmanTree(forest)
	fmt.Println(tree[0].Root.E)
	LevelTraversal(tree[0].Root)
	fmt.Println(c)
}

//贪心算法
//思想：每次从森林中选两个权值最小的树合并，直到森林中只有一颗树
func BuildHuffmanTree(f binaryforest) binaryforest {
	for len(f) != 1 { //直到森林中只有一颗树,这颗树就是哈夫曼树
		f = Sort(f, 0, len(f)-1)
		newtree := merge(f[0], f[1]) //排序后森林的第一颗树和第二课树就是权值最小的树，合并产生新的树，加入森林
		f = append(f[2:], newtree)   //删除两小添新人
	}
	return f
}

//快排:大的后边，小的放前边，以base划分递归
func Sort(f binaryforest, left, right int) binaryforest {
	if left >= right {
		return f
	}
	i, j := left, right
	baseNode := f[i]
	base := f[i].Root.E.(int)
	if left >= right {
		return f
	}

	for i != j {
		for ; i < j && f[j].Root.E.(int) >= base; j-- {
		}
		if i < j {
			f[i] = f[j]
		}
		for ; i < j && f[i].Root.E.(int) <= base; i++ {
		}
		if i < j {
			f[j] = f[i]
		}
	}
	f[i] = baseNode
	Sort(f, 0, i-1)
	Sort(f, i+1, right)
	return f
}

//合并两个树
func merge(a, b *tree.BinaryTree) *tree.BinaryTree {
	tree := tree.NewBinaryTree().InitBinaryTree(a.Root.E.(int) + b.Root.E.(int))
	tree.Root.Left = a.Root
	tree.Root.Right = b.Root
	return tree
}

func LevelTraversal(n *tree.BinaryNode) {
	q := myQueue.NewQueue()
	q.Push(n)
	for !q.NullQueue() {
		p := q.Pop().(*tree.BinaryNode)
		c = append(c, p.E.(int))
		if p.Left != nil {
			q.Push(p.Left)
		}
		if p.Right != nil {
			q.Push(p.Right)
		}

	}
}
