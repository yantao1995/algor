package leetcode

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	nodes := [][]*TreeNode{{root}}
	for i := 0; i < len(nodes); i++ {
		temp := []*TreeNode{}
		for j := 0; j < len(nodes[i]); j++ {
			if nodes[i][j].Left != nil {
				temp = append(temp, nodes[i][j].Left)
			}
			if nodes[i][j].Right != nil {
				temp = append(temp, nodes[i][j].Right)
			}
		}
		if len(temp) > 0 {
			nodes = append(nodes, temp)
		}
	}

	result := make([][]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		result[i] = make([]int, len(nodes[i]))
		for j := 0; j < len(nodes[i]); j++ {
			result[i][j] = nodes[i][j].Val
		}
	}
	return result
}

// @lc code=end
