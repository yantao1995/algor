package leetcode

/*
 * @lc app=leetcode.cn id=1302 lang=golang
 *
 * [1302] 层数最深叶子节点的和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	depthSum := map[int]int{}
	maxDepth := 0
	var recursion func(root *TreeNode, depth int)
	recursion = func(root *TreeNode, depth int) {
		if root != nil {
			depthSum[depth] += root.Val
			if depth > maxDepth {
				maxDepth = depth
			}
			recursion(root.Left, depth+1)
			recursion(root.Right, depth+1)
		}
	}
	recursion(root, 1)
	return depthSum[maxDepth]
}

// @lc code=end

/*
	递归遍历每一层
	depth记录层数，maxDepth 记录最大层数,depthSum记录每一层的和
	然后返回最大层的和即可
*/
