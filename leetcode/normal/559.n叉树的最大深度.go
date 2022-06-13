package leetcode

/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */ //

func maxDepth(root *Node) int {
	max559 := 0
	mid559(root, 0, &max559)
	return max559
}

func mid559(root *Node, depth int, max559 *int) {
	if root != nil {
		depth++
		if depth > *max559 {
			*max559 = depth
		}
		for k := range root.Children {
			mid559(root.Children[k], depth, max559)
		}
	}
}

// @lc code=end

type Node struct {
	Val      int
	Children []*Node
}
