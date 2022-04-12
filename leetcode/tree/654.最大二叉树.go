package leetcode

/*
 * @lc app=leetcode.cn id=654 lang=golang
 *
 * [654] 最大二叉树
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
func constructMaximumBinaryTree(nums []int) *TreeNode {
	var build func(s, e int) *TreeNode
	build = func(s, e int) *TreeNode {
		if s <= e {
			max := nums[s]
			mid := s
			for i := s; i <= e; i++ {
				if max < nums[i] {
					max = nums[i]
					mid = i
				}
			}
			return &TreeNode{
				Val:   nums[mid],
				Left:  build(s, mid-1),
				Right: build(mid+1, e),
			}
		}
		return nil
	}
	return build(0, len(nums)-1)
}

// @lc code=end

/*
	build 记录当前的索引区间
	mid用于在索引区间找到当前的最大值
	用mid构建结点，然后 [s,mid-1] 构建左子树  [mid+1,e]构建右子树
*/
