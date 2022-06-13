package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func isUnivalTree(root *TreeNode) bool {
	val := 0
	if root != nil {
		val = root.Val
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		nodeElement := queue.Back()
		queue.Remove(nodeElement)
		node := nodeElement.Value.(*TreeNode)
		if node != nil {
			if node.Val != val {
				return false
			}
			queue.PushBack(node.Left)
			queue.PushBack(node.Right)
		}
	}
	return true
}

// @lc code=end
