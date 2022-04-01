package leetcode

/*
 * @lc app=leetcode.cn id=437 lang=golang
 *
 * [437] 路径总和 III
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
func pathSum(root *TreeNode, targetSum int) int {
	result := 0
	var treeDp func(node *TreeNode, current int)
	var treeDp2 func(node *TreeNode)
	treeDp = func(node *TreeNode, current int) {
		if node == nil {
			return
		}
		if current+node.Val == targetSum {
			result++
		}
		treeDp(node.Left, current+node.Val)
		treeDp(node.Right, current+node.Val)
	}
	treeDp2 = func(node *TreeNode) {
		if node == nil {
			return
		}
		treeDp(node, 0)
		treeDp2(node.Left)
		treeDp2(node.Right)
	}
	treeDp2(root)
	return result
}

// @lc code=end

/*
	treeDp2  依次向下遍历每个节点，具体先中后都可以
	然后每个节点调用 treeDp 从每个节点开始，依次计算当前节点向下遍历的累加值
	因为节点中包含 负数和0 ，所以在找到等值时，依然继续向下遍历
*/
