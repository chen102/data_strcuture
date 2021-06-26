package main

import (
	"data_structure/Data/Stack"
	tree "data_structure/Data/expressionTree"
	"data_structure/Data/myQueue"
	"fmt"
)

var (
	c  []byte
	c1 []byte
	c2 []byte
	c4 []byte
)

func main() {
	root := &tree.Node{'-', &tree.Node{'+', &tree.Node{'5', nil, nil}, &tree.Node{'*', &tree.Node{'7', nil, nil}, &tree.Node{'3', nil, nil}}}, &tree.Node{'*', &tree.Node{'5', nil, nil}, &tree.Node{'+', &tree.Node{'2', nil, nil}, &tree.Node{'1', nil, nil}}}}
	//inorder_traversal1(root)
	//fmt.Printf("%s\n", c)
	//preorder_traversal(root)
	//fmt.Printf("%s\n", c1)

	//postorder_traversal(root)
	//fmt.Printf("%s\n", c2)

	//Root := NewExpressionTree(c2)
	//fmt.Printf("%c", Root.Left.Right.Left.E)

	preorder_traversal1(root)
	fmt.Printf("%s\n", c1)
	//root1 := &tree.Node{'A',
	//&tree.Node{'B', &tree.Node{E: 'D'}, &tree.Node{E: 'E'}}, &tree.Node{'C', &tree.Node{'F', &tree.Node{E: 'O'}, &tree.Node{E: 'P'}}, &tree.Node{E: 'G'}}}
	//LevelTraversal(root1)
	//fmt.Printf("%s\n", c4)
	//r := copy(root1)
	//LevelTraversal(r)
	//fmt.Printf("%s\n", c4)

	//fmt.Println(deep(r))
	//fmt.Println(sum(r))
	//fmt.Println(LeafNode(r))
}

//表达式树的转换
//前序遍历 根左右
func inorder_traversal(n *tree.Node) {
	c = append(c, n.E)
	if n.Left != nil {
		inorder_traversal(n.Left)
	}
	if n.Right != nil {
		inorder_traversal(n.Right)
	}
}
func inorder_traversal1(n *tree.Node) {
	c = append(c, n.E)
	if n.Left == nil || n.Right == nil {
		return
	}
	inorder_traversal1(n.Left)
	inorder_traversal1(n.Right)
}

//中序遍历 左根右
func preorder_traversal(n *tree.Node) {
	if n.Left != nil {
		preorder_traversal(n.Left)
	}
	c1 = append(c1, n.E)
	if n.Right != nil {
		preorder_traversal(n.Right)
	}

}

//后序遍历 左右根
func postorder_traversal(n *tree.Node) {
	if n.Left != nil {
		postorder_traversal(n.Left)
	}
	if n.Right != nil {
		postorder_traversal(n.Right)
	}

	c2 = append(c2, n.E)

}

//3种遍历算法访问路径都是相同的，只是访问节点时机不同(每个节点经过3次)
//第一次经过：先序遍历  第二次经过:中序遍历 第三次经过:后序遍历
//时间复杂度：O(n) 每个节点访问一次
//空间复杂度:O(n) 栈占用最大辅助空间
//构造一颗表达树
func NewExpressionTree(c []byte) *tree.Node {
	stack := Stack.NewStack1()
	for _, v := range c {
		if v == '+' || v == '-' || v == '*' || v == '/' {
			node := &tree.Node{E: v}
			node.Right = stack.Pop().(*tree.Node)
			node.Left = stack.Pop().(*tree.Node)
			stack.Push(node)
		} else {
			stack.Push(&tree.Node{E: v})
		}
	}
	return stack.Pop().(*tree.Node)
}

//不用递归中序遍历(栈) 也是递归的原理
func preorder_traversal1(n *tree.Node) {
	stack := Stack.NewStack1()
	for n != nil || !stack.NullStack() {
		if n != nil {
			stack.Push(n) //栈里存指针,就可以回溯了
			n = n.Left
		} else {
			p := stack.Pop().(*tree.Node)
			c1 = append(c1, p.E)
			n = p.Right
		}
	}
}

//层次遍历(队列)
func LevelTraversal(n *tree.Node) {
	q := myQueue.NewQueue()
	q.Push(n)
	for !q.NullQueue() {
		p := q.Pop().(*tree.Node)
		c4 = append(c4, p.E)
		if p.Left != nil {
			q.Push(p.Left)
		}
		if p.Right != nil {
			q.Push(p.Right)
		}

	}
}

//二叉树的应用(递归应用)
//复制二叉树
func copy(oldn *tree.Node) (newn *tree.Node) {
	if oldn == nil {
		return nil
	} else {
		newn := &tree.Node{E: oldn.E}
		newn.Left = copy(oldn.Left)
		newn.Right = copy(oldn.Right)
		return newn
	}
}

//计算二叉树的深度
func deep(n *tree.Node) int {
	if n == nil {
		return 0
	} else {
		m := deep(n.Left)  //计算左子树的深度
		n := deep(n.Right) //计算右子树的深度
		if m > n {
			return m + 1
		} else {
			return n + 1
		}
	}
}

//计算二叉树节点总数
func sum(n *tree.Node) int {
	if n == nil {
		return 0
	} else {
		return sum(n.Left) + sum(n.Right) + 1
	}
}

//计算叶子节点的个数(没有孩子的节点)
func LeafNode(n *tree.Node) int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		return 1
	} else {
		return LeafNode(n.Left) + LeafNode(n.Right)
	}
}
