package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=655 lang=golang
 *
 * [655] 输出二叉树
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
func printTree(root *TreeNode) [][]string {
	height := 0
	var getHeight func(root *TreeNode, currentHeight int)
	getHeight = func(root *TreeNode, currentHeight int) {
		if root != nil {
			currentHeight++
			if currentHeight > height {
				height = currentHeight
			}
			getHeight(root.Left, currentHeight)
			getHeight(root.Right, currentHeight)
		}
	}
	getHeight(root, 0)
	result := make([][]string, height)
	for k := range result {
		result[k] = make([]string, 1<<height-1)
	}
	var writeResult func(root *TreeNode, rangeStart, rangeEnd, currentHeight int)
	writeResult = func(root *TreeNode, rangeStart, rangeEnd, currentHeight int) {
		if root != nil {
			midIndex := (rangeStart + rangeEnd) / 2
			result[currentHeight][midIndex] = strconv.Itoa(root.Val)
			writeResult(root.Left, rangeStart, midIndex-1, currentHeight+1)
			writeResult(root.Right, midIndex+1, rangeEnd, currentHeight+1)
		}
	}
	writeResult(root, 0, 1<<height-1, 0)
	return result
}

// @lc code=end

/*
	递归遍历层，写入树结点值
	height 记录当前树高度  1-n
	getHeight() 拿到当前的树高度
	result 使用 height 生成当前树形状数组
	writeResult()根据当前的结点以及当前结点的左右区间范围，取中间值，写入当前结点值
*/
