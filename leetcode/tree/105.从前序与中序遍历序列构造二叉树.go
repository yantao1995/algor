package leetcode

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(pIdx *int, inl, inr int) *TreeNode
	build = func(pIdx *int, inl, inr int) *TreeNode {
		if inr >= inl && *pIdx < len(preorder) {
			nodeInIdx := 0
			for i := inl; i <= inr; i++ {
				if inorder[i] == preorder[*pIdx] {
					nodeInIdx = i
					break
				}
			}
			temp := *pIdx
			*pIdx++
			return &TreeNode{
				Val:   preorder[temp], //递归指针最后赋值
				Left:  build(pIdx, inl, nodeInIdx-1),
				Right: build(pIdx, nodeInIdx+1, inr),
			}
		}
		return nil
	}
	pIdx := 0
	return build(&pIdx, 0, len(preorder)-1)
}

// @lc code=end
