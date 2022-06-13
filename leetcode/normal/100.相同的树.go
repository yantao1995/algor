package leetcode

/*
 * @lc app=leetcode.cn id=100 lang=golang
 *
 * [100] 相同的树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */ //
func isSameTree(p *TreeNode, q *TreeNode) bool {
	pi, qi := &[]interface{}{}, &[]interface{}{}
	first(p, pi)
	first(q, qi)
	if len(*pi) != len(*qi) {
		return false
	}
	pp, qq := *pi, *qi
	for i := 0; i < len(pp); i++ {
		if pp[i] != qq[i] {
			return false
		}
	}
	return true
}
func first(t *TreeNode, i *[]interface{}) {
	if t != nil {
		*i = append(*i, t.Val)
		first(t.Left, i)
		first(t.Right, i)
	} else {
		*i = append(*i, nil)
	}
}

// @lc code=end

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
