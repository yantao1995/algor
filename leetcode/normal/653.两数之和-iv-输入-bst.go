package leetcode

/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func findTarget(root *TreeNode, k int) bool {
	data := []int{}
	mid653(root, &data)
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == k {
				return true
			}
		}
	}
	return false
}

func mid653(root *TreeNode, data *[]int) {
	if root != nil {
		mid653(root.Left, data)
		*data = append(*data, root.Val)
		mid653(root.Right, data)
	}
}

// @lc code=end
