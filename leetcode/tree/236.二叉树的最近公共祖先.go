package leetcode

/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root != nil {
		ln := lowestCommonAncestor(root.Left, p, q)
		rn := lowestCommonAncestor(root.Right, p, q)
		//fmt.Println(root.Val, ln, rn)
		if root == p || root == q {
			return root
		}
		if ln != nil && rn != nil {
			return root
		} else if ln != nil {
			return ln
		} else if rn != nil {
			return rn
		}

	}
	return nil
}

// @lc code=end
