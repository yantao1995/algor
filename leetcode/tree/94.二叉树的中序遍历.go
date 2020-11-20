package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */
//

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	lst := list.New()
	for lst.Len() > 0 || root != nil {
		if root != nil {
			lst.PushBack(root)
			root = root.Left
		} else {
			node := lst.Back()
			result = append(result, node.Value.(*TreeNode).Val)
			root = node.Value.(*TreeNode).Right
			lst.Remove(node)
		}
	}
	return result
}

// @lc code=end
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
