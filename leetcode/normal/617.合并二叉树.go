package leetcode

/*
 * @lc app=leetcode.cn id=617 lang=golang
 *
 * [617] 合并二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	firstMerge617(t1, t1, t2, true)
	return t1
}

func firstMerge617(t1, t1Prev, t2 *TreeNode, isleft bool) {
	if t1 != nil || t2 != nil {
		if t1 != nil && t2 != nil {
			if &t1 != &t2 {
				t1.Val += t2.Val
			}
			firstMerge617(t1.Left, t1, t2.Left, true)
			firstMerge617(t1.Right, t1, t2.Right, false)
		}
		if t1 == nil && t2 != nil {
			if isleft {
				t1Prev.Left = t2
			} else {
				t1Prev.Right = t2
			}
		}
	} else {
		t1 = t2
	}

}

// @lc code=end
