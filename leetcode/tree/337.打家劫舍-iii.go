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

/*
	当前节点有两种状态，偷或者不偷，那就取偷或者不偷状态下，向下取所有子节点的最大值
	初始时：dpTrue, dpFalse 分别标识root节点的初始偷还是不偷的状态
	如果当前节点可以偷，那么左右子节点就是false状态，即不可偷状态，直接取左右节点只和
	如果当前节点不可偷，那么左右节点可以选择投或者不偷，即4种状态  左：偷  不偷  右：偷  不偷 的排列组合
*/
