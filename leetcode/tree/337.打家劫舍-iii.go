package leetcode

/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
func rob(root *TreeNode) int {
	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dpTrue, dpFalse := map[*TreeNode]int{}, map[*TreeNode]int{}
	var dfs func(root *TreeNode, isRob bool) int
	dfs = func(root *TreeNode, isRob bool) int {
		if root == nil {
			return 0
		}
		if isRob {
			if val, ok := dpTrue[root]; ok {
				return val
			}
			val := root.Val + dfs(root.Left, false) + dfs(root.Right, false)
			dpTrue[root] = val
			return val
		} else {
			if val, ok := dpFalse[root]; ok {
				return val
			}
			val := getMax(
				getMax(dfs(root.Left, true)+dfs(root.Right, false), dfs(root.Left, true)+dfs(root.Right, true)),
				getMax(dfs(root.Left, false)+dfs(root.Right, true), dfs(root.Left, false)+dfs(root.Right, false)))
			dpFalse[root] = val
			return val
		}
	}
	return getMax(dfs(root, true), dfs(root, false))
}

// @lc code=end
