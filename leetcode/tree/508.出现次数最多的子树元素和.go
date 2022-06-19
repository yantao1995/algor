package leetcode

/*
 * @lc app=leetcode.cn id=508 lang=golang
 *
 * [508] 出现次数最多的子树元素和
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
func findFrequentTreeSum(root *TreeNode) []int {
	mc := map[int]int{}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxCount := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		root.Val += dfs(root.Left) + dfs(root.Right)
		mc[root.Val]++
		maxCount = max(maxCount, mc[root.Val])
		return root.Val
	}
	dfs(root)
	result := []int{}
	for k := range mc {
		if mc[k] == maxCount {
			result = append(result, k)
		}
	}
	return result
}

// @lc code=end

/*
	使用mc记录每个节点的累加值
	使用maxCount记录最大的累加值
	然后遍历结果集，把最大值加进去
*/
