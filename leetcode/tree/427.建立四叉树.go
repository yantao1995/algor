package leetcode

/*
 * @lc app=leetcode.cn id=427 lang=golang
 *
 * [427] 建立四叉树
 */

// @lc code=start
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

// func construct(grid [][]int) *Node {
// 	isLeaf := func(is, ie, js, je int) (int, bool) { //矩阵的范围坐标 是不是叶子
// 		if is < 0 || js < 0 || ie >= len(grid) || je >= len(grid) {
// 			return 1, true
// 		}
// 		currentVal := grid[is][js]
// 		for i := is; i <= ie; i++ {
// 			for j := js; j <= je; j++ {
// 				if grid[i][j] != currentVal {
// 					return 0, false
// 				}
// 			}
// 		}
// 		return currentVal, true
// 	}
// 	var createNode func(is, ie, js, je int) *Node
// 	createNode = func(is, ie, js, je int) *Node {
// 		if is <= ie && js <= je {
// 			if val, ok := isLeaf(is, ie, js, je); ok {
// 				return &Node{
// 					Val:         val == 1,
// 					IsLeaf:      true,
// 					TopLeft:     nil,
// 					TopRight:    nil,
// 					BottomLeft:  nil,
// 					BottomRight: nil,
// 				}
// 			}
// 			il, jl := (ie-is+1)/2, (je-js+1)/2
// 			return &Node{
// 				Val:         false,
// 				IsLeaf:      false,
// 				TopLeft:     createNode(is, is+il-1, js, js+jl-1),
// 				TopRight:    createNode(is, is+il-1, js+jl, je),
// 				BottomLeft:  createNode(is+il, ie, js, js+jl-1),
// 				BottomRight: createNode(is+il, ie, js+jl, je),
// 			}
// 		}
// 		return nil
// 	}
// 	return createNode(0, len(grid)-1, 0, len(grid)-1)
// }

// // @lc code=end

// type Node struct {
// 	Val         bool
// 	IsLeaf      bool
// 	TopLeft     *Node
// 	TopRight    *Node
// 	BottomLeft  *Node
// 	BottomRight *Node
// }

/*
	题目描述 ::: 翻译
	如果当前网格的值相同（即，全为 0 或者全为 1），将 isLeaf 设为 True ，将 val 设为网格相应的值，并将四个子节点都设为 Null 然后停止。
		：：：当前网格值相同，那么本结点就是叶子结点。
	如果当前网格的值不同，将 isLeaf 设为 False， 将 val 设为任意值，然后如下图所示，将当前网格划分为四个子网格。
		：：：不同，那么就不是叶子结点
	使用适当的子网格递归每个子节点。
		：：：根据图中描述，上面的两个是 top 的l和r 下面的 bottom 的l 和r


	判断当前is,ie,js,je 标识的矩阵是否为叶子结点，
	如果不是叶子结点，就划分成四等分
	然后依次递归
*/
