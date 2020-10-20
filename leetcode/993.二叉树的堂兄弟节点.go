package leetcode

import (
	"container/list"
)

/*
 * @lc app=leetcode.cn id=993 lang=golang
 *
 * [993] 二叉树的堂兄弟节点
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func isCousins(root *TreeNode, x int, y int) bool {
	lst := list.New()
	lstNew := list.New()
	lstNew.PushFront(root)
	xFlag, yFlag := false, false
	index, xIndex, yIndex := 0, 0, 0
	xyFlag := 0 //1x  2y
	for lstNew.Len() > 0 {
		lst = lstNew
		lstNew = list.New()
		index = 0
		for lst.Len() > 0 {
			index++
			element := lst.Back()
			lst.Remove(element)
			node := element.Value.(*TreeNode)
			if node != nil {
				if node.Val == x || node.Val == y {
					if node.Val == x {
						xFlag = true
						xIndex = index
						xyFlag = 1
					} else {
						yFlag = true
						yIndex = index
						xyFlag = 2
					}
					goto comeHere
				}
				lstNew.PushFront(node.Left)
				lstNew.PushFront(node.Right)
			}
		}
	}
comeHere:
	for lst.Len() > 0 {
		index++
		element := lst.Back()
		lst.Remove(element)
		node := element.Value.(*TreeNode)
		if node != nil {
			if node.Val == x {
				xFlag = true
				xIndex = index
			}
			if node.Val == y {
				yFlag = true
				yIndex = index
			}
		}
		if xFlag && yFlag {
			if xyFlag == 1 {
				if xIndex%2 == 1 {
					if yIndex-xIndex == 1 {
						return false
					}
				}
			} else {
				if yIndex%2 == 1 {
					if xIndex-yIndex == 1 {
						return false
					}
				}
			}
		}
	}
	return xFlag == true && yFlag == true
}

// @lc code=end
