package leetcode

/*
 * @lc app=leetcode.cn id=530 lang=golang
 *
 * [530] 二叉搜索树的最小绝对差
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func getMinimumDifference(root *TreeNode) int {
	min := 1<<63 - 1
	first530(root, &min)
	return min
}

func first530(root *TreeNode, minValue *int) {
	if root != nil {
		if root.Left != nil {
			nodeL := root.Left
			for nodeL != nil {
				if root.Val-nodeL.Val < *minValue {
					*minValue = root.Val - nodeL.Val
				}
				nodeL = nodeL.Right
			}
			first530(root.Left, minValue)
		}
		if root.Right != nil {
			nodeR := root.Right
			for nodeR != nil {
				if nodeR.Val-root.Val < *minValue {
					*minValue = nodeR.Val - root.Val
				}
				nodeR = nodeR.Left
			}
			first530(root.Right, minValue)
		}
	}
}

// @lc code=end
