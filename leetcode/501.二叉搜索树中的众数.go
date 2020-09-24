package leetcode

/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func findMode(root *TreeNode) []int {
	result := []int{}
	ht := map[int]int{}
	mid501(root, ht)
	count := 0
	for k := range ht {
		if ht[k] > count {
			count = ht[k]
		}
	}
	for k := range ht {
		if ht[k] == count {
			result = append(result, k)
		}
	}
	return result
}
func mid501(root *TreeNode, ht map[int]int) {
	if root != nil {
		ht[root.Val]++
		mid501(root.Left, ht)
		mid501(root.Right, ht)
	}
}

// @lc code=end
