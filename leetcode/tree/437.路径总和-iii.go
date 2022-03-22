package leetcode

/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	result := 0
	var treeDp func(node *TreeNode, current int)
	var treeDp2 func(node *TreeNode)
	treeDp = func(node *TreeNode, current int) {
		if node == nil {
			return
		}
		if current+node.Val == targetSum {
			result++
		}
		treeDp(node.Left, current+node.Val)
		treeDp(node.Right, current+node.Val)
	}
	treeDp2 = func(node *TreeNode) {
		if node == nil {
			return
		}
		treeDp(node, 0)
		treeDp2(node.Left)
		treeDp2(node.Right)
	}
	treeDp2(root)
	return result
}

// @lc code=end
