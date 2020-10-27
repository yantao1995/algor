package easy

import "strconv"

/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] 根据二叉树创建字符串
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func tree2str(t *TreeNode) string {
	result := ""
	first606(t, &result)
	return result
}

func first606(root *TreeNode, str *string) {
	if root != nil {
		*str += strconv.Itoa(root.Val)
		if root.Left != nil || root.Right != nil {
			*str += "("
			first606(root.Left, str)
			*str += ")"
			if root.Right != nil {
				*str += "("
				first606(root.Right, str)
				*str += ")"
			}
		}
	}
}

// @lc code=end
