package leetcode

/*
 * @lc app=leetcode.cn id=222 lang=golang
 *
 * [222] 完全二叉树的节点个数
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
func countNodes(root *TreeNode) int {
	if root != nil {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
	return 0
}

// @lc code=end

/*
	递归直接数每一个
*/
