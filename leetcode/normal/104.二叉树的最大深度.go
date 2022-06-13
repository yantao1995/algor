package leetcode

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */ //
// func maxDepth(root *TreeNode) int {
// 	max104 := 0
// 	mid104(root, 0, &max104)
// 	return max104
// }

// func mid104(root *TreeNode, depth int, max104 *int) {
// 	if root != nil {
// 		depth++
// 		if depth > *max104 {
// 			*max104 = depth
// 		}
// 		mid104(root.Left, depth, max104)
// 		mid104(root.Right, depth, max104)
// 	}
// }

// @lc code=end
