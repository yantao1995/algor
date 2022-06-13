package leetcode

/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func minDiffInBST(root *TreeNode) int {
	min := 1<<63 - 1
	first783(root, &min)
	return min
}

func first783(root *TreeNode, minValue *int) {
	if root != nil {
		if root.Left != nil {
			nodeL := root.Left
			for nodeL != nil {
				if root.Val-nodeL.Val < *minValue {
					*minValue = root.Val - nodeL.Val
				}
				nodeL = nodeL.Right
			}
			first783(root.Left, minValue)
		}
		if root.Right != nil {
			nodeR := root.Right
			for nodeR != nil {
				if nodeR.Val-root.Val < *minValue {
					*minValue = nodeR.Val - root.Val
				}
				nodeR = nodeR.Left
			}
			first783(root.Right, minValue)
		}
	}
}

// @lc code=end
