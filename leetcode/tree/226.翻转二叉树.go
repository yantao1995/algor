package leetcode

/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
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
func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	}
	return root
}

// @lc code=end

/*
	直接左右交叉赋值即可
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	或者 :
	left := invertTree(root.Left)
	right:= invertTree(root.Right)
	root.Right, root.Left = left, right

	不能
	root.Right = invertTree(root.Left)
	root.Left = invertTree(root.Right)
	因为 第一步的交换已经影响了第二步的交换了，两边都同时处理的 原始的root.left
*/
