package leetcode

/*
 * @lc app=leetcode.cn id=938 lang=golang
 *
 * [938] 二叉搜索树的范围和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func rangeSumBST(root *TreeNode, L int, R int) int {
	result := 0
	first938(root, &L, &R, &result)
	return result
}

func first938(root *TreeNode, L, R, result *int) {
	if root != nil {
		if root.Val >= *L && root.Val <= *R {
			*result += root.Val
		}
		first938(root.Left, L, R, result)
		first938(root.Right, L, R, result)
	}
}

// @lc code=end
