package leetcode

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */ //

func preorder(root *Node) []int {
	result := []int{}
	first589(root, &result)
	return result
}

func first589(root *Node, nums *[]int) {
	if root != nil {
		*nums = append(*nums, root.Val)
		for k := range root.Children {
			first589(root.Children[k], nums)
		}
	}
}

// @lc code=end
