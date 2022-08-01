package leetcode

/*
 * @lc app=leetcode.cn id=1161 lang=golang
 *
 * [1161] 最大层内元素和
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
func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root, nil}
	result := 1
	maxLevel := root.Val
	for i, temp, level := 0, 0, 1; i < len(queue); i++ {
		if queue[i] == nil {
			if queue[i-1] != nil {
				if temp > maxLevel {
					maxLevel = temp
					result = level
				}
				temp = 0
				level++
				queue = append(queue, nil)
			}
		} else {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			temp += queue[i].Val
		}
	}
	return result
}

// @lc code=end

/*
	用队列来层序遍历，使用nil来做层分隔
	temp做层累加，level记录当前层号，maxLevel记录最大层的值
	遇到queue[i]不为nil就累加，为nil就做层处理
*/
