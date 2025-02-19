package lcpr

/*
 * @lc app=leetcode.cn id=1367 lang=golang
 * @lcpr version=30204
 *
 * [1367] 二叉树中的链表
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	flag := false
	var find func(h *ListNode, r *TreeNode) bool
	find = func(h *ListNode, r *TreeNode) bool {
		if flag {
			return true
		}
		if h == nil || r == nil {
			return h == nil && r == nil
		}
		if h.Val == r.Val {
			if find(h.Next, r.Left) || find(h.Next, r.Right) {
				flag = true
				return flag
			}
		}
		return find(h, r.Left) || find(h, r.Right)
	}
	find(head, root)
	return flag
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,8]\n[1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,4,2,6]\n[1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [1,4,2,6,8]\n[1,4,4,null,2,2,null,1,null,6,8,null,null,null,null,1,3]\n
// @lcpr case=end

*/
