package leetcode

/*
 * @lc app=leetcode.cn id=998 lang=golang
 *
 * [998] 最大二叉树 II
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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root != nil {
		if root.Val < val {
			return &TreeNode{
				Val:   val,
				Left:  root,
				Right: nil,
			}
		}
		root.Right = insertIntoMaxTree(root.Right, val)
		return root
	}
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

// @lc code=end

/*
	因为 val 被添加在末尾，所以如果val结点被添加在树的非叶子结点时，
	那么它的所有子节点都应该是左孩子结点下的结点。

	val结点的3个位置：
		1.头结点
		2.右子树向右递归的某个节点
		3.子节点

	1和2可以合并。即当前结点不为空时，没有走到末尾。
		那么当前结点如果小于 val，那么成为val的左孩子结点
		val成为当前结点的原父节点的子节点
	3。走到末尾仍然没有比val结点小的结点，那么val就应该是右下角叶子
		结点的右孩子结点
*/
