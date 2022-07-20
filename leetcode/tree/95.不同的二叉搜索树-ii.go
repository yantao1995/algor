package leetcode

/*
 * @lc app=leetcode.cn id=95 lang=golang
 *
 * [95] 不同的二叉搜索树 II
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
func generateTrees(n int) []*TreeNode {
	var generate func(left, right int) []*TreeNode
	generate = func(left, right int) []*TreeNode {
		result := []*TreeNode{}
		for i := left; i <= right; i++ {
			ls := generate(left, i-1)
			rs := generate(i+1, right)
			if len(ls) == 0 && len(rs) == 0 {
				result = append(result, &TreeNode{i, nil, nil})
			} else if len(ls) == 0 {
				for _, r := range rs {
					result = append(result, &TreeNode{i, nil, r})
				}
			} else if len(rs) == 0 {
				for _, l := range ls {
					result = append(result, &TreeNode{i, l, nil})
				}
			} else {
				for _, l := range ls {
					for _, r := range rs {
						result = append(result, &TreeNode{i, l, r})
					}
				}
			}
		}
		return result
	}
	return generate(1, n)
}

// @lc code=end

/*
	递归生成左右子树,
	因为是二叉搜索树，需要保证左子树小于右子树，
	递归的每个结点，都在当前层当一次根结点，
	然后小于当前值的，都为左子树，大于当前值的，都为右子树
	所以从left -> right 的过程中：
		left - i-1   是左子树
		i+1  - right 是右子树
	这样必然满足当前结点值小于左子树，大于右子树
	然后递归保证每颗子树都满足条件
*/
