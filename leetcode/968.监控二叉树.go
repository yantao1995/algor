package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=968 lang=golang
 *
 * [968] 监控二叉树
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
func minCameraCover(root *TreeNode) int {
	stack := []*TreeNode{root, nil} //trees
	count := []int{}                // 每层多少个结点
	for i := 0; i < len(stack); i++ {
		if stack[i] != nil {
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		} else {
			if i == len(stack)-1 {
				break
			}
			stack = append(stack, nil)
		}
	}
	current := 0
	for i := range stack {
		if stack[i] != nil {
			current++
		} else {
			count = append(count, current)
			current = 0
		}
	}
	fmt.Println(count)

	return 0
}

// @lc code=end
