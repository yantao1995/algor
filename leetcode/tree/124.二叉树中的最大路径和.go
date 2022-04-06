package leetcode

/*
 * @lc app=leetcode.cn id=124 lang=golang
 *
 * [124] 二叉树中的最大路径和
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
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	mp := -1000
	var maxPathInner func(root *TreeNode) int
	maxPathInner = func(root *TreeNode) int {
		l, r := 0, 0
		mp = max(root.Val, mp)
		if root.Left != nil {
			l = maxPathInner(root.Left)
			mp = max(l, mp)
		}
		if root.Right != nil {
			r = maxPathInner(root.Right)
			mp = max(r, mp)
		}
		mp = max(root.Val+l+r, mp)
		mp = max(mp, root.Val+max(0, max(l, r)))
		return root.Val + max(0, max(l, r))
	}
	maxPathInner(root)
	return mp
}

// @lc code=end

/*
	maxPathInner 用于选择当前结点，或者当前结点往下的一条边
		___2___
	   -1      3
	比例上述结点，判断当前结点左右路径是否最大，有下面几种选择：
	1.如果要选择当前的根节点且包含子节点，即选择2。则结果要么 -1 -> 2 -> 3  要么  -1 ->2  要么 2->3
	2.如果不选择根节点，那左右 -1 和 3 选其一即可。
	3.如果只选择父节点，则结果是 2 。

	然后记录下当前选择路径的最大值，再向上层结点传递包含本节点及下层节点的最大值

*/
