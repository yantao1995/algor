package leetcode

/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
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
func longestUnivaluePath(root *TreeNode) int {
	maxLength := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root != nil {
			left := dfs(root.Left)
			right := dfs(root.Right)
			if root.Left != nil && root.Val == root.Left.Val {
				left++
			} else {
				left = 0
			}
			if root.Right != nil && root.Val == root.Right.Val {
				right++
			} else {
				right = 0
			}
			maxLength = max(maxLength, left+right)
			return max(left, right)
		}
		return 0
	}
	dfs(root)
	return maxLength
}

// @lc code=end

/*
	两点构成一条边，求最长的一条线

	后续遍历，先拿到左右的最长线，然后判断能否通过root连接起来，
	如果能连接，那么判断是否是最长。
	如果不能连接，就将对应边置0
	最后返回加上左右之后的最长的一条线
*/
