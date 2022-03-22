package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	result := 0
	distinct := map[string]bool{}
	var treeDp func(node *TreeNode, current int, dist string)
	treeDp = func(node *TreeNode, current int, dist string) {
		if node == nil {
			return
		}
		//fmt.Println(current, node.Val, current+node.Val == targetSum, !distinct[dist+strconv.Itoa(node.Val)])
		if current+node.Val == targetSum && !distinct[dist+strconv.Itoa(node.Val)] {
			result++
			distinct[dist+strconv.Itoa(node.Val)] = true
		}
		treeDp(node.Left, 0, "l"+strconv.Itoa(node.Val))
		treeDp(node.Right, 0, "r"+strconv.Itoa(node.Val))
		treeDp(node.Left, current+node.Val, "l"+dist+strconv.Itoa(node.Val))
		treeDp(node.Right, current+node.Val, "r"+dist+strconv.Itoa(node.Val))
	}
	if root != nil {
		treeDp(root, 0, "rt"+strconv.Itoa(root.Val))
	}
	return result
}

// @lc code=end
