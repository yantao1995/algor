package leetcode

/*
 * @lc app=leetcode.cn id=623 lang=golang
 *
 * [623] 在二叉树中增加一行
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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}

	var dfs func(node *TreeNode, currentDepth int)
	dfs = func(node *TreeNode, currentDepth int) {
		if node != nil {
			if currentDepth > 1 {
				dfs(node.Left, currentDepth-1)
				dfs(node.Right, currentDepth-1)
			} else {
				node.Left = &TreeNode{val, node.Left, nil}
				node.Right = &TreeNode{val, nil, node.Right}
			}
		}
	}
	dfs(root, depth-1)
	return root
}

// @lc code=end

/*
	特殊处理root，如果depth为1，那么直接重建结点返回
	对于之后的结点，做dfs处理，使用currentDepth记录当前结点的深度
	找到需要插入行的父节点，然后对应生成子节点即可，然后把父节点的左右
	孩子结点分别挂在新生成的结点上
*/
