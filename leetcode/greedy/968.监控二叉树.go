package leetcode

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
	head := root
	result := 0
	var recursion func(root *TreeNode) int
	recursion = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		l := recursion(root.Left)
		r := recursion(root.Right)
		if l == 2 && r == 2 {
			if root == head {
				result++
			}
			return 0
		}
		if l == 1 || r == 1 {
			if l == 0 || r == 0 {
				result++
				return 1
			}
			return 2
		}
		result++
		return 1
	}
	recursion(root)
	return result
}

// @lc code=end

/*
	每次贪子节点的父节点为摄像头，因为至少可以覆盖两层，
	就算父节点没有爷爷结点，那么也可以覆盖2层，有的话，就是3层。
	而如果选择子节点，那么最好的情况才能覆盖 2层。
	0表示未覆盖，1表示摄像头，2表示覆盖
	底层递归子节点判断：左右无叶子，相当于被覆盖了，就不用管，那么自身就是0
	上层递归父节点判断，如果左右有一个没有被覆盖，那么自己就应该成为摄像头来覆盖下层
	上层递归爷爷结点判断，如果左右有一个为摄像头，那么自身就是被覆盖，但是如果再有一个为未覆盖，
		那么自身应该成为摄像头来覆盖。其余情况均应该成为摄像头
*/
