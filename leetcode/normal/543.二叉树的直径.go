package leetcode

/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * } */ //
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	first543(root, &maxDiameter)
	return maxDiameter
}

func first543(root *TreeNode, maxDiameter *int) { //不一定经过根节点。每个结点依次遍历
	if root != nil {
		leftRadius, rightRadius := 0, 0
		if root.Left != nil {
			left543(root.Left, 1, &leftRadius)
		}
		if root.Right != nil {
			right543(root.Right, 1, &rightRadius)
		}
		if leftRadius+rightRadius > *maxDiameter {
			*maxDiameter = leftRadius + rightRadius
		}
		first543(root.Left, maxDiameter)
		first543(root.Right, maxDiameter)
	}
}

func left543(root *TreeNode, currentRadius int, leftRadius *int) {
	if root != nil {
		if currentRadius > *leftRadius {
			*leftRadius = currentRadius
		}
		left543(root.Left, currentRadius+1, leftRadius)
		left543(root.Right, currentRadius+1, leftRadius)
	}
}

func right543(root *TreeNode, currentRadius int, rightRadius *int) {
	if root != nil {
		if currentRadius > *rightRadius {
			*rightRadius = currentRadius
		}
		right543(root.Left, currentRadius+1, rightRadius)
		right543(root.Right, currentRadius+1, rightRadius)
	}
}

// @lc code=end
