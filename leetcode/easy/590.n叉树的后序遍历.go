package leetcode

/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */ //

func postorder(root *Node) []int {
	result := []int{}
	last589(root, &result)
	return result
}

func last589(root *Node, nums *[]int) {
	if root != nil {
		for k := range root.Children {
			last589(root.Children[k], nums)
		}
		*nums = append(*nums, root.Val)
	}
}

// @lc code=end
