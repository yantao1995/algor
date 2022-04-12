package leetcode

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	flag := true
	var recursion func(root *TreeNode, depth int) int
	recursion = func(root *TreeNode, depth int) int {
		if root != nil {
			depth++
			if root.Left != nil && root.Right != nil {
				l := recursion(root.Left, depth)
				r := recursion(root.Right, depth)
				if l > r {
					if l-r > 1 {
						flag = false
					}
					return l
				}
				if r-l > 1 {
					flag = false
				}
				return r
			} else if root.Left != nil {
				max := recursion(root.Left, depth)
				if max-depth > 1 {
					flag = false
				}
				return max
			} else if root.Right != nil {
				max := recursion(root.Right, depth)
				if max-depth > 1 {
					flag = false
				}
				return max
			}
		}
		return depth
	}
	recursion(root, 0)
	return flag
}

// @lc code=end

/*
	用dept记录当前层数
	每个结点都判断左右高度是否差距大于1
	左右结点都存在的情况下，返回最长的路径
*/
