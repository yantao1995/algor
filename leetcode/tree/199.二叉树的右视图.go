package leetcode

import (
	"container/list"
)

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	nodes := []int{}
	q1, q2 := list.New(), list.New()
	if root == nil {
		return nodes
	}
	q1.PushBack(root)
	for q1.Len() > 0 || q2.Len() > 0 {
		if element := q1.Front(); element != nil {
			if node, ok := element.Value.(*TreeNode); ok {
				if node.Left != nil {
					q2.PushBack(node.Left)
				}
				if node.Right != nil {
					q2.PushBack(node.Right)
				}
				if q1.Len() == 1 {
					nodes = append(nodes, node.Val)
				}
			}
			q1.Remove(element)
		} else if q2.Len() > 0 {
			for q2e := q2.Front(); q2e != nil; q2e = q2e.Next() {
				q1.PushBack(q2e.Value.(*TreeNode))
			}
			q2.Init()
		} else {
			break
		}
	}
	return nodes
}

// @lc code=end
