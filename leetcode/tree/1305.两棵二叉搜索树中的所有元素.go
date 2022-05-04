package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1305 lang=golang
 *
 * [1305] 两棵二叉搜索树中的所有元素
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
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	result := []int{}
	var recursion func(root *TreeNode)
	recursion = func(root *TreeNode) {
		if root != nil {
			recursion(root.Left)
			result = append(result, root.Val)
			recursion(root.Right)
		}
	}
	recursion(root1)
	recursion(root2)
	sort.Ints(result)
	return result
}

// @lc code=end

/*
读取所有的值，然后排序
*/
