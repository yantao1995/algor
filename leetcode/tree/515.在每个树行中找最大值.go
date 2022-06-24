package leetcode

/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) []int {
	mMax := map[int]int{}
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root != nil {
			if v, ok := mMax[level]; !ok || root.Val > v {
				mMax[level] = root.Val
			}
			dfs(root.Left, level+1)
			dfs(root.Right, level+1)
		}
	}
	dfs(root, 0)
	result := make([]int, len(mMax))
	for k := range result {
		result[k] = mMax[k]
	}
	return result
}

// @lc code=end

/*
	mMax 记录每层的最大值
	dfs遍历每一层，然后依次比较
	注意：结果集需要按层数递增顺序返回
*/
