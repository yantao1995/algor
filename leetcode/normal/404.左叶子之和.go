package leetcode

/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	sum404(root, &sum)
	return sum
}

func sum404(root *TreeNode, sum *int) int {
	if root != nil {
		if root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				*sum += root.Left.Val
			}
		}
		sum404(root.Left, sum)
		sum404(root.Right, sum)
	}
	return *sum
}

// @lc code=end
