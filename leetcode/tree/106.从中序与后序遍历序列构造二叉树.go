package leetcode

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
// func buildTree(inorder []int, postorder []int) *TreeNode {
// 	var build func(is, ie int, pi *int) *TreeNode
// 	build = func(is, ie int, pi *int) *TreeNode {
// 		if is <= ie {
// 			mid := 0
// 			for i := is; i <= ie; i++ {
// 				if inorder[i] == postorder[*pi] {
// 					mid = i
// 					break
// 				}
// 			}
// 			*pi--
// 			t := &TreeNode{
// 				Val:   inorder[mid],
// 				Right: build(mid+1, ie, pi),
// 				Left:  build(is, mid-1, pi),
// 			}
// 			return t
// 		}
// 		return nil
// 	}
// 	pi := len(postorder) - 1
// 	return build(0, len(inorder)-1, &pi)
// }

// @lc code=end

/*
	与 105 的 前序+中序 构造二叉树差不多
	但是需要注意的是，生成树左右结点的顺序，因为是 中序+后序，所以需要先生成右结点
	按后序倒序的结点顺序来递归划分中序数组的左右边界
*/
