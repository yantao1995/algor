package leetcode

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
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
func minDepth(root *TreeNode) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	var recursion func(root *TreeNode, depth int) int
	recursion = func(root *TreeNode, depth int) int {
		if root != nil {
			depth++
			if root.Left != nil && root.Right != nil {
				return min(recursion(root.Left, depth),
					recursion(root.Right, depth))
			} else if root.Left != nil {
				return recursion(root.Left, depth)
			} else if root.Right != nil {
				return recursion(root.Right, depth)
			}
		}
		return depth
	}
	return recursion(root, 0)
}

// @lc code=end

/*
	递归向下
	如果左右结点都不为空，就取两个的最小值,否则进入单边子结点
	都为空，说明当前层级就是末尾层级，直接返回
*/
