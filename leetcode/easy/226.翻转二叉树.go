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
 * } */ //
func invertTree(root *TreeNode) *TreeNode {
	res226(root)
	return root
}
func res226(root *TreeNode) {
	if root != nil {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp
		res226(root.Left)
		res226(root.Right)
	}
}

// @lc code=end
