package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=671 lang=golang
 *
 * [671] 二叉树中第二小的节点
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
func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	rootVal := root.Val
	childMinVal := -1
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		element := queue.Front()
		if node, ok := element.Value.(*TreeNode); ok && node != nil {
			if node.Val == rootVal {
				queue.PushBack(node.Left)
				queue.PushBack(node.Right)
			} else if childMinVal > node.Val || childMinVal == -1 {
				childMinVal = node.Val
			}
		}
		queue.Remove(element)
	}
	return childMinVal
}

// @lc code=end
