package leetcode

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	result := 0
	var recursion func(root *TreeNode)
	recursion = func(root *TreeNode) {
		if root != nil {
			if k < 0 {
				return
			}
			recursion(root.Left)
			k--
			if k == 0 {
				result = root.Val
			}
			recursion(root.Right)
		}
	}
	recursion(root)
	return result
}

// @lc code=end

/*
	因为是二叉搜索树，索引需要中序遍历才能得到递增顺序的序列
	recursion 做递归中序遍历
	result 记录第k个值。
*/
