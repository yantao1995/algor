package graph

import (
	"algor/queue"
	"algor/stack"
	"algor/vals"
	"fmt"
)

const (
	MaxEdge = 5     //结点数
	Unable  = -9999 //不可达值
)

//有向图
type Graph struct {
	EdgesWeight [MaxEdge][MaxEdge]int   //边权值
	Visited     [MaxEdge]bool           //遍历标志 初始默认 flase
	VexValue    [MaxEdge]vals.AlgorType //结点值
	VertexNum   int                     //结点数
	EdgeNum     int                     //边数
}

//BFS 广度
func (gph Graph) BFS(start int) {
	q := queue.Queue{}
	q.LPush(start) // 初始结点入队列
	fmt.Printf("%v\t", gph.VexValue[start])
	gph.Visited[start] = true
	for q.Size > 0 {
		index := q.RPop()
		for j := start; j < MaxEdge; j++ {
			if !gph.Visited[j] && gph.EdgesWeight[index.(int)][j] != Unable {
				fmt.Printf("%v\t", gph.VexValue[j])
				q.LPush(j)
				gph.Visited[j] = true
			}
		}
	}
}

//DFS 深度
func (gph Graph) DFS(start int) {
	s := stack.Stack{}
	s.Push(start) // 初始结点入栈
	gph.Visited[start] = true
	fmt.Printf("%v\t", gph.VexValue[start])
	for s.Size > 0 {
		index := s.GetTopButNoPop() // 获取当前结点
		for i := start; i < MaxEdge; i++ {
			if !gph.Visited[i] && gph.EdgesWeight[index.(int)][i] != Unable {
				s.Push(i)
				fmt.Printf("%v\t", gph.VexValue[i])
				gph.Visited[i] = true
				break
			}
			if i == MaxEdge-1 { // 当前结点下无子结点，出栈
				s.Pop()
			}
		}
	}
}

//Dijkstra 单源最短路径   带权有向图
func (gph Graph) Dijkstra(start, end int) {
	dist := [MaxEdge]int{} //初始化 路径参数
	for i := start; i < MaxEdge; i++ {
		dist[i] = gph.EdgesWeight[start][i]
	}
	for i := 0; i < MaxEdge; i++ { // dijkstra

	}
}

/*		结构
		      A
	B<————12——|——————10——> C______
	|		  ⬇25				  |
  12|————————>D———20———>E<——————13
*/
//TestGraph 测试
func TestGraph() {
	gph := InitGraph()
	start := 0 //  小于 MaxEdge
	println("start:", start, "广度BFS:")
	gph.BFS(start)
	println("\nstart:", start, "深度DFS:")
	gph.DFS(start)
	println("\nDijstra:")
}

/*
type Graph struct {
	EdgesWeight [MaxEdge][MaxEdge]int   //边权值
	Visited     [MaxEdge]bool            //遍历标志 初始默认 flase
	VexValue    [MaxEdge]vals.AlgorType //结点值
	VertexNum   int                     //结点数
	EdgeNum     int                     //边数
}
*/
//InitGraph 初始化创建
func InitGraph() Graph {
	graph := Graph{}
	graph.VertexNum = MaxEdge
	graph.EdgeNum = 6
	graph.VexValue[0] = "A"
	graph.VexValue[1] = "B"
	graph.VexValue[2] = "C"
	graph.VexValue[3] = "D"
	graph.VexValue[4] = "E"
	for i := 0; i < MaxEdge; i++ { //初始化邻接矩阵
		for j := 0; j < MaxEdge; j++ {
			graph.EdgesWeight[i][j] = Unable
		}
	}
	graph.EdgesWeight[0][1] = 12
	graph.EdgesWeight[0][2] = 10
	graph.EdgesWeight[0][3] = 25
	graph.EdgesWeight[1][3] = 12
	graph.EdgesWeight[3][4] = 20
	graph.EdgesWeight[2][4] = 35
	for i := 0; i < MaxEdge; i++ { //输出邻接矩阵
		for j := 0; j < MaxEdge; j++ {
			fmt.Printf("\t%v", graph.EdgesWeight[i][j])
		}
		println()
	}
	return graph
}
