package leetcode

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := -1 << 63
	first543(root, 0, &maxDiameter)
	return maxDiameter
}

func first543(root *TreeNode, diameter int, maxDiameter *int) {
	if root != nil {

	}
}

// @lc code=end
