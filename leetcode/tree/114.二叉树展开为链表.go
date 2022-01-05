package leetcode

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	nums := []*TreeNode{}
	first114(&nums, root)
	for i := 1; i < len(nums); i++ {
		root.Right = nums[i]
		root.Left = nil
		root = root.Right
	}
}

func first114(nums *[]*TreeNode, root *TreeNode) {
	if root != nil {
		*nums = append(*nums, root)
		first114(nums, root.Left)
		first114(nums, root.Right)
	}
}

// @lc code=end
