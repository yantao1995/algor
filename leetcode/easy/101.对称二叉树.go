package leetcode

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	pi, qi := &[]interface{}{}, &[]interface{}{}
	firstL(root.Left, pi)
	firstR(root.Right, qi)
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
func firstL(t *TreeNode, i *[]interface{}) {
	if t != nil {
		*i = append(*i, t.Val)
		firstL(t.Left, i)
		firstL(t.Right, i)
	} else {
		*i = append(*i, nil)
	}
}
func firstR(t *TreeNode, i *[]interface{}) {
	if t != nil {
		*i = append(*i, t.Val)
		firstR(t.Right, i)
		firstR(t.Left, i)
	} else {
		*i = append(*i, nil)
	}
}

// @lc code=end
