package leetcode

/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //

// 解析 ：https://github.com/echofoo/ARTS/blob/master/A-20190925-%E6%9C%80%E9%95%BF%E5%90%8C%E5%80%BC%E8%B7%AF%E5%BE%84.md
func longestUnivaluePath(root *TreeNode) int {
	max := 0
	recursion687(root, &max)
	return max
}

func recursion687(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := recursion687(root.Left, max)
	right := recursion687(root.Right, max)
	leftMax, rightMax := 0, 0
	if root.Left != nil && root.Left.Val == root.Val {
		leftMax = left + 1
	}
	if root.Right != nil && root.Right.Val == root.Val {
		rightMax = right + 1
	}
	if leftMax+rightMax > *max {
		*max = leftMax + rightMax
	}
	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}

// @lc code=end
