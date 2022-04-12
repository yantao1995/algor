package leetcode

/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root != nil {
// 		if root == q || root == p {
// 			return root
// 		}
// 		l := lowestCommonAncestor(root.Left, p, q)
// 		r := lowestCommonAncestor(root.Right, p, q)
// 		if l != nil && r != nil {
// 			return root
// 		} else if l != nil {
// 			return l
// 		} else if r != nil {
// 			return r
// 		}
// 	}
// 	return nil
// }

// @lc code=end

/*
	递归判断当前结点是不是p,q，如果是就直接返回
	如果p,q结点在同一条线上，那么直接就找到了父结点
	如果不在一条线上，那么分别返回l和r，就找到了当前的左右结点
	如果左右结点都找到了，那么当前结点就是父结点
*/
