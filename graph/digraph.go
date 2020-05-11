package graph

import (
	"algor/vals"
	"fmt"
)

const (
	MaxEdge = 5    //结点数
	Unable  = -999 //不可达值
)

//有向图
type Graph struct {
	EdgesWeight [MaxEdge][MaxEdge]int   //边权值
	Visited     [MaxEdge]int            //遍历标志 初始默认0
	VexValue    [MaxEdge]vals.AlgorType //结点值
	VertexNum   int                     //结点数
	EdgeNum     int                     //边数
}

//BFS 广度
func BFS() {

}

//DFS 深度
func DFS() {

}

//Dijkstra 单源最短路径   带权有向图
func Dijkstra() {

}

/*		结构
		      A
	B<————12——|——————10——> C______
	|		  ⬇25				  |
  12|————————>D———20———>E<——————13
*/
//TestGraph 测试
func TestGraph() {
	// graph := InitGraph()
	// fmt.Println(graph)
	s := [1][1]int{}
	check(&s)
	fmt.Println(s)
}

/*
type Graph struct {
	EdgesWeight [MaxEdge][MaxEdge]int   //边权值
	Visited     [MaxEdge]int            //遍历标志 初始默认0
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
	// for i := 0; i < MaxEdge; i++ { //输出邻接矩阵
	// 	for j := 0; j < MaxEdge; j++ {
	// 		fmt.Printf("\t%v", graph.Edges[i][j])
	// 	}
	// 	println()
	// }
	return graph
}

func check(s *[1][1]int) {
	s[0][0] = 1
}
