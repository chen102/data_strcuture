package main

import (
	"data_structure/Data/MGrape"
	"data_structure/Data/myQueue"
	"fmt"
)

func main() {
	//grape, err := MGrape.CreateMGraph3(6, 6)
	//if err != nil {
	//fmt.Println(err)
	//return
	//}
	//RangeGrape(grape, make([]bool, len(grape.VertexType)))
	grape, err := MGrape.CreateMGraph3(9, 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	RangeGrape(grape, make([]bool, len(grape.VertexType)), false)
}

//O(n*n)
//只能对连通分量进行遍历,借助rangegrape遍历多个连通分量
func DFS(grape *MGrape.MGrape, start int, visited []bool) []bool {
	fmt.Println(grape.VertexType[start].(string))
	visited[start] = true
	//fmt.Println(visited)
	for i := 0; i < len(grape.VertexType); i++ {
		if grape.EdgeType[start][i] == 1 && visited[i] == false {
			DFS(grape, i, visited)
		}
	}
	return visited
}

func BFS(grape *MGrape.MGrape, start int, visited []bool) []bool {
	q := myQueue.NewQueue()
	q.Push(start)
	visited[start] = true
	for !q.NullQueue() {
		p := q.Pop().(int)
		fmt.Println(grape.VertexType[p])
		for i := 0; i < len(grape.VertexType); i++ {
			if grape.EdgeType[p][i] == 1 && visited[i] == false {
				q.Push(i)
				visited[i] = true
			}
		}
	}
	return visited
}
func RangeGrape(grape *MGrape.MGrape, visited []bool, DFSORBFS bool) {
	var i int
	for k, v := range visited {
		if v == false {
			if DFSORBFS { //若DFSORBFS=true为DFS遍历
				visited = DFS(grape, k, visited) //刷新visited  //每次选取visited中第一个为false的节点
				i++
			} else { //若DFSORBFS=false为BFS遍历
				visited = BFS(grape, k, visited)
				i++
			}
		}
	}
	fmt.Println(i, "个连通分量")
}

//遍历连通树的结果为该图的生成树(深度优先生成树、广度优先生成树)
//生成树:n个节点n-1条边
//最小生成树:给定一个无向网络,在该网络中的所有生成树中，使得各边权值之和最小的那颗生成树
//MST性质
//普利姆算法O(n*n) 找点
//克鲁斯卡尔O(e*loge) 找边
//dijkstra(n*n)
//思想：每次将权值最小的节点加入，更新最开始节点到所有节点的距离
func dijkstra(grape *MGrape.MGrape) []int {
	var V []int
	var p int
	var visited []bool
	for k, _ := range grape.VertexType {
		V[k] = grape.EdgeType[0][k]
	}
	visited[0] = true
	for range grape.VertexType {
		min := 32767
		for k, v := range V {
			if V[k] < min && !visited[k] {
				min = v
				p = k
			}
			visited[p] = true
			for key, _ := range grape.VertexType {
				if grape.EdgeType[p][key] < 32767 {
					if V[p]+grape.EdgeType[p][key] < V[key] {
						V[key] = V[p] + grape.EdgeType[p][key]
					}
				}
			}
		}
	}
	return V
}
