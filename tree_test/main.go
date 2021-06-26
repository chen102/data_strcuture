package main

import "data_structure/Data/tree"

import "fmt"

func main() {
	f := tree.NewForest(5)
	f = f.NewTree("A")
	n := f[0].Root.AddChild("B")
	n1 := n.AddBrother("L")
	n1 = n1.AddChild("H")
	n = n.AddChild("C")
	n = n.AddBrother("D")
	n = n.AddBrother("E")
	//tree.RangeTree(f[0].Root)
	//fmt.Println("-----------------")
	//tree.RangeTree1(f[0].Root)
	//fmt.Println("-----------------")
	//tree.RangeTree2(f[0].Root)
	binarytree := tree.NewBinaryTree()
	binarytreeRoot := tree.TreetoBinaryTree(f[0].Root, binarytree.Root)
	RangeBinaryTree(binarytreeRoot) //遍历二叉树
	fmt.Println("根节点的值为:", binarytreeRoot.Left.Tree.Root.E)

	f = f.InitTree()
	btree := tree.BinaryNodetoTree(binarytreeRoot, f[1].Root)
	tree.RangeTree(btree)
	fmt.Println("深度为:", btree.Deep())
	//fmt.Println(btree.Child().Child().Brother().Element().(string)) //验证一下
	//构建一个森林
	forest := tree.NewForest(3)
	forest = forest.NewTree("A")
	forest = forest.NewTree("E")
	forest = forest.NewTree("G")
	forest[0].Root.AddChild("B").AddBrother("C").AddBrother("D")
	forest[1].Root.AddChild("F")
	forest[2].Root.AddChild("H").AddBrother("I").AddChild("J")

	binarytree1 := tree.NewBinaryTree()
	binarytreeRoot1 := tree.ForesttoBinaryTree(forest, binarytree1.Root)
	fmt.Println(binarytreeRoot1)
	RangeBinaryTree(binarytreeRoot1)
	fmt.Println(binarytreeRoot1.Right.Right.Left.Right) //验证一下

	forest1 := tree.BinaryTreetoForest(binarytreeRoot1)
	fmt.Println("----------------------------------------")
	fmt.Println(forest1[0].Root)
	fmt.Println(forest1[1].Root)
	fmt.Println(forest1[2].Root)
	fmt.Println(forest1[2].Root.Child().Brother().Child())

	fmt.Println("----------------------------------------")
	tree.RangeForest(forest1)

}
func RangeBinaryTree(n *tree.BinaryNode) {
	if n == nil {
		return
	} else {
		fmt.Println(n.E)
		RangeBinaryTree(n.Left)
		RangeBinaryTree(n.Right)
	}
}
