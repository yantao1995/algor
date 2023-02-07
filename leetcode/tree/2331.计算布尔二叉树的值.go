package leetcode

/*
 * @lc app=leetcode.cn id=2331 lang=golang
 *
 * [2331] 计算布尔二叉树的值
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
func evaluateTree(root *TreeNode) bool {
	if root.Left != nil && root.Right != nil {
		l := evaluateTree(root.Left)
		r := evaluateTree(root.Right)
		if root.Val == 2 {
			return l || r
		} else {
			return l && r
		}
	}
	return root.Val == 1
}

// @lc code=end

/*
	直接递归即可
*/
