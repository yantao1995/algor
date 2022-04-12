package leetcode

/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	var search func(root *TreeNode)
	search = func(root *TreeNode) {
		if root.Val > val {
			if root.Left == nil {
				root.Left = node
				return
			}
			search(root.Left)
		} else {
			if root.Right == nil {
				root.Right = node
				return
			}
			search(root.Right)
		}
	}
	if root == nil {
		return node
	}
	search(root)
	return root
}

// @lc code=end

/*
	判断当前结点与val的大小关系，然后分别进入左右子树
	如果要进入的左右子树不存在，那么就可以构建结点了
*/
