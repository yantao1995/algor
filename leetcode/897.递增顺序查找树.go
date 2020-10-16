package leetcode

/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序查找树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func increasingBST(root *TreeNode) *TreeNode {
	new := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	middle897(root, new)
	return new.Right
}

func middle897(root, new *TreeNode) {
	if root != nil {
		middle897(root.Left, new)
		for new.Right != nil {
			new = new.Right
		}
		new.Right = &TreeNode{
			Val:   root.Val,
			Left:  nil,
			Right: nil,
		}
		middle897(root.Right, new)
	}
}

// @lc code=end
