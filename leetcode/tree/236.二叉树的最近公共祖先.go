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

/*
	后序遍历
	直接递归找，如果左结点和右结点都有，那最近公共祖先就是自己
	如果只有左右单节点是自己，那就直接返回自己
	如果左右单节点存在，那就返回左右单节点
*/
