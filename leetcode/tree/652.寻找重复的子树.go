package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	count := map[string]int{}
	result := []*TreeNode{}
	var dfs func(root *TreeNode) string
	dfs = func(root *TreeNode) string {
		if root != nil {
			ser := fmt.Sprintf("%d[%s][%s]", root.Val, dfs(root.Left), dfs(root.Right))
			count[ser]++
			if count[ser] == 2 {
				result = append(result, root)
			}
			return ser
		}
		return ""
	}
	dfs(root)
	return result
}

// @lc code=end

/*
	后序遍历，自下而上
	序列化树结构，然后放到记录相同的序列化结构出现的次数，
	等于2的就加入到结果集中

	------------------------

	最开始想直接暴力枚举子树，但是发现去重会超时

	func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	logs := []*TreeNode{}
	eqMap := map[*TreeNode]bool{}
	result := []*TreeNode{}
	var equals func(r1, r2 *TreeNode) bool
	equals = func(r1, r2 *TreeNode) bool {
		if r1 != nil && r2 != nil {
			return r1.Val == r2.Val &&
				equals(r1.Left, r2.Left) &&
				equals(r1.Right, r2.Right)
		}
		return r1 == nil && r2 == nil
	}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			for k := range logs {
				if equals(logs[k], root) {
					if !eqMap[logs[k]] {
						result = append(result, root)
					}
					eqMap[logs[k]] = true
					eqMap[root] = true
				}
			}
			logs = append(logs, root)
			//dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(root)
	return result
}
*/
