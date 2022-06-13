package leetcode

/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	rightFirst538(root, &sum)
	return root
}

func rightFirst538(root *TreeNode, sum *int) {
	if root != nil {
		rightFirst538(root.Right, sum)
		root.Val += *sum
		*sum = root.Val
		rightFirst538(root.Left, sum)
	}
}

// @lc code=end
