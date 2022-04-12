package leetcode

/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
func findBottomLeftValue(root *TreeNode) int {
	result := 0
	maxDepth := 0
	var recursion func(root *TreeNode, depth int)
	recursion = func(root *TreeNode, depth int) {
		if root != nil {
			depth++
			if depth > maxDepth {
				maxDepth = depth
				result = root.Val
			}
			recursion(root.Left, depth)
			recursion(root.Right, depth)
		}
	}
	recursion(root, 0)
	return result
}

// @lc code=end

/*
	本质上就是找最高的那个结点
	depth记录每层的高度，然后依次比较判断
*/
