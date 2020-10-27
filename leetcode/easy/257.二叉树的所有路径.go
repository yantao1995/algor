package easy

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func binaryTreePaths(root *TreeNode) []string {
	sli := []string{}
	search257(root, "", &sli)
	return sli
}

func search257(root *TreeNode, s string, sli *[]string) {
	if root != nil {
		if s == "" {
			s += strconv.Itoa(root.Val)
		} else {
			s += "->" + strconv.Itoa(root.Val)
		}
		if root.Left == nil && root.Right == nil {
			*sli = append(*sli, s)
		}
		search257(root.Left, s, sli)
		search257(root.Right, s, sli)
	}
}

// @lc code=end
