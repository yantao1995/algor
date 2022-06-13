package leetcode

/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] 叶子相似的树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	slice1, slice2 := []int{}, []int{}
	first872(root1, &slice1)
	first872(root2, &slice2)
	if len(slice1) != len(slice2) {
		return false
	}
	for k := range slice1 {
		if slice1[k] != slice2[k] {
			return false
		}
	}
	return true
}

func first872(root *TreeNode, slice *[]int) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			*slice = append(*slice, root.Val)
		}
		first872(root.Left, slice)
		first872(root.Right, slice)
	}
}

// @lc code=end
