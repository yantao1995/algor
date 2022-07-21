package leetcode

/*
 * @lc app=leetcode.cn id=814 lang=golang
 *
 * [814] 二叉树剪枝
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
func pruneTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left = pruneTree(root.Left)
		root.Right = pruneTree(root.Right)
		if root.Left == nil &&
			root.Right == nil &&
			root.Val == 0 {
			return nil
		}
	}
	return root
}

// @lc code=end

/*

	后序遍历

	只要子树有一个是1，就会层层向上反馈影响，从而当前路径不需要删除。
	所以需要优先处理左右子树是否需要删除，如果左右子树
	都删除了或者一开始就是nil，再考虑当前结点是否需要删除。





//优化前的版本
func pruneTree(root *TreeNode) *TreeNode {
	var recursion func(root *TreeNode) bool
	recursion = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		l := recursion(root.Left)
		r := recursion(root.Right)
		if !l {
			root.Left = nil
		}
		if !r {
			root.Right = nil
		}
		return l || r || root.Val == 1
	}
	if !recursion(root) {
		root = nil
	}
	return root
}
*/
