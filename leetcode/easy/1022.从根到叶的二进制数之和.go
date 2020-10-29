package leetcode

/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func sumRootToLeaf(root *TreeNode) int {
	result, parent := 0, 0
	first1022(root, parent, &result)
	return result
}

func first1022(root *TreeNode, parent int, result *int) {
	if root != nil {
		parent *= 2
		parent += root.Val
		if root.Left == nil && root.Right == nil {
			*result += parent
		}
		first1022(root.Left, parent, result)
		first1022(root.Right, parent, result)
	}
}

// @lc code=end
