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

/*
	preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	前序遍历的顺序为当前根节点，因为是第一个节点，所有左右结点的域为 [0,len(preorder)-1]
	如当前遍历preorder 到3,则在 inorder 中找到3，
	此时 inorder 中3左边[0,0]的所有值为左子树，右边[2,4]的为右子树
	构建左子树：此时遍历preorder 到9,则在 inorder 中找到9，
	此时 inorder 中9左边[0,-1]的所有值为左子树，右边[1,0]的为右子树，
	依次类推.....
*/
