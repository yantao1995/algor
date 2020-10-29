package leetcode

/*
 * @lc app=leetcode.cn id=563 lang=golang
 *
 * [563] 二叉树的坡度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func findTilt(root *TreeNode) int {
	result := 0
	last563(root, &result)
	return result
}

func last563(root *TreeNode, result *int) {
	if root != nil {
		last563(root.Left, result)
		last563(root.Right, result)
		root.Val += getSum563(root.Left, root.Right)
		if root.Left != nil || root.Right != nil {
			*result += getAbsoult563(root.Left, root.Right)
		}
	}
}
func getSum563(a, b *TreeNode) int {
	if a != nil && b != nil {
		return a.Val + b.Val
	}
	if a == nil && b != nil {
		return b.Val
	}
	if a != nil && b == nil {
		return a.Val
	}
	return 0
}

func getAbsoult563(a, b *TreeNode) int {
	if a != nil && b != nil {
		if a.Val > b.Val {
			return a.Val - b.Val
		}
		return b.Val - a.Val
	}
	if a == nil && b != nil {
		if b.Val < 0 {
			return 0 - b.Val
		}
		return b.Val
	}
	if b == nil && a != nil {
		if a.Val < 0 {
			return 0 - a.Val
		}
		return a.Val
	}
	return 0
}

// @lc code=end
