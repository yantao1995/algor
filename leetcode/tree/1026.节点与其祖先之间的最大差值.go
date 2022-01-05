package leetcode

/*
 * @lc app=leetcode.cn id=1026 lang=golang
 *
 * [1026] 节点与其祖先之间的最大差值
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
func maxAncestorDiff(root *TreeNode) int {
	maxVal := 0
	temp := 0
	if root == nil {
		return maxVal
	}
	first1026(root, root.Val, root.Val, &maxVal, &temp)
	return maxVal
}
func first1026(root *TreeNode, min, max int, maxVal, temp *int) {
	if root != nil {
		if root.Val > max {
			max = root.Val
		}
		if root.Val < min {
			min = root.Val
		}
		getMaxVal1026(maxVal, temp, min, max)
		first1026(root.Left, min, max, maxVal, temp)
		first1026(root.Right, min, max, maxVal, temp)
	}
}
func getMaxVal1026(maxVal, temp *int, val1, val2 int) {
	*temp = val1 - val2
	if *temp < 0 {
		*temp = 0 - *temp
	}
	if *temp > *maxVal {
		*maxVal = *temp
	}
}

// @lc code=end
