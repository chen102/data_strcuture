package main

import (
	"bytes"
	tree "data_structure/Data/Threaded_Binary_Tree"
	"encoding/binary"
	"fmt"
)

var c []byte
var Node []*tree.BinartNode

func main() {
	//binaryTreeRoot := tree.NewBinaryTree().InitTreeRoot('A') //'A':int32
	//fmt.Println(binaryTreeRoot)
	//n := binaryTreeRoot.AddLeftNode('B')
	//n.AddLeftNode('C')
	//n2 := n.AddRightNode('D')
	//n3 := n2.AddLeftNode('E')
	//n2.AddRightNode('F')
	//n3.AddRightNode('G')

	//RangeTree(binaryTreeRoot)
	//fmt.Printf("%s\n", c)
	//GetRangeNode(binaryTreeRoot)

	//ThreadedBinartTree := Clueing(binaryTreeRoot)
	//fmt.Println("---------------------------")
	//fmt.Printf("%c", ThreadedBinartTree.Left.Right.Right.Right.E)

	//再实验一下
	Node = nil
	binaryTreeRoot1 := tree.NewBinaryTree().InitTreeRoot('A')
	binaryTreeRoot1.AddLeftNode('B').AddRightNode('C')
	binaryTreeRoot1.AddRightNode('D').AddLeftNode('E')
	//c = nil
	//RangeTree(binaryTreeRoot1)
	//fmt.Printf("%s\n", c)
	//GetRangeNode(binaryTreeRoot1)
	//ThreadedBinartTree1 := Clueing(binaryTreeRoot1)
	//IsNull(ThreadedBinartTree1.Left.Right.Right)
	pre = binaryTreeRoot1

	ThreadedBinartTree2 := Clueing1(binaryTreeRoot1)
	//RangeThreadedBinaryTree(ThreadedBinartTree2)
	RangeThreadedBinaryTree1(ThreadedBinartTree2)
}
func RangeTree(n *tree.BinartNode) {
	if n == nil {
		return
	} else {
		RangeTree(n.Left)
		s := IntToByte(n.E.(int32))
		c = append(c, s[len(s)-1]) //大端，取最后一个元素
		RangeTree(n.Right)
	}
}
func GetRangeNode(n *tree.BinartNode) {
	if n == nil {
		return
	} else {
		GetRangeNode(n.Left)
		Node = append(Node, n) //直接存储节点地址
		GetRangeNode(n.Right)
	}

}
func IntToByte(n int32) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n) //大端模式:高位字节排在内存的低地址端，低位字节排放在内存的高地址端 小端模式：相反
	return bytesBuffer.Bytes()
}

var i = -1 //Node的指针
//在递归中遍历Node，由于Node也是中序遍历拿到的，所以是一一对应
//线索化（中序） 先序和后序一样的
//思想:先获得中序遍历得到的节点Node,然后再中序遍历，每次递归处理的节点保证和i指向Node的节点一致,就可以直接为那个节点添加前驱和后继
func Clueing(n *tree.BinartNode) *tree.BinartNode {
	if n == nil {
		return nil
	} else {
		Clueing(n.Left) //只要这里不接收返回值，就不会影响递归程序,最后递归接收了返回设定的值
		i++             //此时的Node[i]应该和正在处理的节点相同
		//fmt.Printf("n:%c\n", n.E)
		//fmt.Println("i:", i)
		//fmt.Println("Node:", Node[i])
		if n.Left == nil && i != 0 { //添加前驱节点，确保Node[i]前面有值
			n.Left = Node[i-1]
			//fmt.Printf("n.Left:%c\n", Node[i-1].E)

			n.Ltag = tree.Thread
		}
		if n.Right == nil && i != len(Node)-1 { //添加后继节点，确保Node[i]后面有值
			n.Right = Node[i+1]
			//fmt.Printf("n.Right:%c\n", Node[i+1].E)
			n.Rtag = tree.Thread
		}
		//上面可能会修改右孩子节点,用Rtag判断是否修改,若修改了，直接返回(若不判断返回，又会进入下一轮递归)
		if n.Rtag == tree.Thread {
			return n
		}
		Clueing(n.Right)
	}
	return n
}
func IsNull(n *tree.BinartNode) {
	if n == nil {
		fmt.Println("是空的")
	} else {
		fmt.Printf("不是空的node:%c\n", n.E)
	}
}

//上面的线索化要两次递归(我自己写的)，下面只要一次(书上的)
var pre *tree.BinartNode //那个全局变量记一下前驱节点不就可以了

func Clueing1(n *tree.BinartNode) *tree.BinartNode {
	if n != nil {
		Clueing1(n.Left)
		if n.Left == nil {
			n.Left = pre
			n.Ltag = tree.Thread
		}
		if pre.Right == nil { //这是我没想到的,在当前节点处理前缀节点的右孩子
			pre.Right = n
			pre.Rtag = tree.Thread
		}
		pre = n
		Clueing1(n.Right)

	}
	return n
}

//这样线索化有个问题，当处理第一个节点时，它将有前驱节点(根)
//解决：Rtag和Ltag不用bool类型了,使用常量类型Thread和Link即tag=Thread说明有前驱或后继节点，tag=Link说明没有前驱和后继节点(也就是处理的第一个节点的前驱和最后一个节点后继)
//获取中序遍历最后一个节点
func InOrderLastNode(n *tree.BinartNode) *tree.BinartNode {
	for n.Rtag == tree.Thread || n.Right != nil {
		n = n.Right
	}
	return n
}
func InOrderFirstNode(n *tree.BinartNode) *tree.BinartNode {
	for n.Ltag == tree.Link {
		n = n.Left
	}
	return n
}

//构建双向链表
func BuildDoubleLink(n *tree.BinartNode) *tree.BinartNode {
	//构建头节点，指向中序遍历处理的第一个节点和最后一个节点
	headnode := &tree.BinartNode{Tree: n.Tree}
	firstnode := InOrderFirstNode(n)
	lastnode := InOrderLastNode(n)
	headnode.Right = n
	headnode.Left = lastnode
	//中序遍历首次处理的节点的后继和前驱指向头节点
	firstnode.Left = headnode
	firstnode.Ltag = tree.Thread //要加线索，否则前序遍历会死循环
	lastnode.Right = headnode
	lastnode.Rtag = tree.Thread
	return headnode
}

//后继遍历线索二叉树
func RangeThreadedBinaryTree(n *tree.BinartNode) {
	headnode := BuildDoubleLink(n)
	for n != headnode {
		for n.Ltag == tree.Link {
			n = n.Left
		}
		fmt.Printf("%c\n", n.E)
		for n.Rtag == tree.Thread && n.Right != headnode {
			n = n.Right
			fmt.Printf("%c\n", n.E)
		}
		n = n.Right

	}

}

//前驱遍历线索二叉树
func RangeThreadedBinaryTree1(n *tree.BinartNode) {
	headnode := BuildDoubleLink(n)
	for n != headnode {
		for n.Rtag == tree.Link {
			n = n.Right
		}
		fmt.Printf("%c\n", n.E)
		for n.Ltag == tree.Thread && n.Left != headnode {
			n = n.Left
			fmt.Printf("%c\n", n.E)
		}
		n = n.Left
	}
}
