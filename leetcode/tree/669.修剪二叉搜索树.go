package leetcode

/*
 * @lc app=leetcode.cn id=669 lang=golang
 *
 * [669] 修剪二叉搜索树
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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root != nil {
		if root.Val < low {
			return trimBST(root.Right, low, high)
		} else {
			root.Left = trimBST(root.Left, low, high)
		}
		if root.Val > high {
			return trimBST(root.Left, low, high)
		} else {
			root.Right = trimBST(root.Right, low, high)
		}
	}
	return root
}

// @lc code=end

/*

	因为是二叉搜索树，左孩子比右孩子值大，所以判断当前节点：
		如果当前节点小于low，那么应该直接剪去当前节点，返回右孩子节点的搜索结果
		如果当前节点大于high，那么应该也剪去当前节点，返回左孩子节点的搜索结果
	（因为左右孩子向下递归也可能超过low与high的范围，所以要向下搜索）
*/
