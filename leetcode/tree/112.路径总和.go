package leetcode

/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root != nil {
		targetSum -= root.Val
		if targetSum == 0 && root.Left == nil && root.Right == nil {
			return true
		}
		return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
	}
	return false
}

// @lc code=end

/*
	targetSum 加法看成减法,走到哪个结点就减哪个结点
	然后判断当前节点的值是不是为0，如果为0，并且没有左右子节点，就返回
*/
