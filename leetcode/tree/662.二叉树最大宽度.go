package leetcode

/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] 二叉树最大宽度
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
func widthOfBinaryTree(root *TreeNode) int {
	levelMin, levelMax := map[int]int{}, map[int]int{}
	maxWidth := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var getLevelLocation func(root *TreeNode, level, rank int)
	getLevelLocation = func(root *TreeNode, level, rank int) {
		if root != nil {
			if levelMin[level] == 0 || rank < levelMin[level] {
				levelMin[level] = rank
			}
			if levelMax[level] == 0 || rank > levelMax[level] {
				levelMax[level] = rank
			}
			getLevelLocation(root.Left, level+1, rank*2-1)
			getLevelLocation(root.Right, level+1, rank*2)
			maxWidth = max(maxWidth, levelMax[level]-levelMin[level]+1)
		}
	}
	getLevelLocation(root, 0, 1)
	return maxWidth
}

// @lc code=end

/*
	因为结构视作满二叉树，所以父节点的层排位，与子节点的层排位关系为
		childLeft =  parent * 2 -1
		childRight =  parent *2
	那么可以存储 levelMin, levelMax 每一层的最小rank和最大rank,然后与maxWidth比较即可
*/
