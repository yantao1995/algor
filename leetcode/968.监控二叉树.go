package leetcode

/*
 * @lc app=leetcode.cn id=968 lang=golang
 *
 * [968] 监控二叉树
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
func minCameraCover(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}
	result := 0
	var recursion func(root *TreeNode, level int)
	recursion = func(root *TreeNode, level int) {
		if root != nil {
			recursion(root.Left, level+1)
			recursion(root.Right, level+1)
		}
	}
	recursion(root, 0)
	return result
}

// @lc code=end
