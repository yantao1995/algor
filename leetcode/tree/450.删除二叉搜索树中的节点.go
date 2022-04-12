package leetcode

/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	deleteNode := func(root *TreeNode) *TreeNode { //右子树接到左子树
		if root.Right != nil {
			if root.Left != nil {
				rTree := root.Right
				lTree := root.Left
				for lTree.Right != nil {
					lTree = lTree.Right
				}
				lTree.Right = rTree
				return root.Left
			}
			return root.Right
		}
		return root.Left
	}

	var deleteFunc func(parent, root *TreeNode, isLeft bool)
	deleteFunc = func(parent, root *TreeNode, isLeft bool) {
		if root != nil {
			if root.Val > key {
				deleteFunc(root, root.Left, true)
			} else if root.Val < key {
				deleteFunc(root, root.Right, false)
			} else {
				if isLeft {
					parent.Left = deleteNode(root)
				} else {
					parent.Right = deleteNode(root)
				}
			}
		}
	}
	if root != nil && root.Val == key {
		l := root.Left
		r := root.Right
		if l == nil {
			return r
		}
		if r == nil {
			return l
		}
		deleteNode(root)
		return l
	}
	deleteFunc(nil, root, true)
	return root
}

// @lc code=end

/*
	头结点单独处理，需要单独判断左右子树是否是nil，
	如果是nil 只需要返回左右子树的下一个结点即可

	中间结点，因为要删除本结点，之后又得连接父节点，所以得记录一个父节点

	具体删除 node 的时候，是直接把左孩子结点接过来，然后把 右孩子结点接到
	左孩子结点的最右结点
	这样就能满足二叉搜索树的性质了
*/
