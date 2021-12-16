package leetcode

/*
 * @lc app=leetcode.cn id=103 lang=golang
 *
 * [103] 二叉树的锯齿形层序遍历
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
func zigzagLevelOrder(root *TreeNode) [][]int {
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
			if i%2 == 0 {
				result[i][j] = nodes[i][j].Val
			} else {
				result[i][len(nodes[i])-j-1] = nodes[i][j].Val
			}

		}
	}
	return result
}

// @lc code=end
