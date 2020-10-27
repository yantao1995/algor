package easy

/*
 * @lc app=leetcode.cn id=669 lang=golang
 *
 * [669] 修剪二叉搜索树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root != nil {
		if root.Val < low {
			return trimBST(root.Right, low, high)
		}
		if root.Val > high {
			return trimBST(root.Left, low, high)
		}
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
		return root
	}
	return nil
}

// @lc code=end
