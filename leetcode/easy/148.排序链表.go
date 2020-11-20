package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//
func sortList(head *ListNode) *ListNode {
	nums := []int{}
	h := head
	for h != nil {
		nums = append(nums, h.Val)
		h = h.Next
	}
	sort.Ints(nums)
	h2 := head
	k := 0
	for h2 != nil {
		h2.Val = nums[k]
		h2 = h2.Next
		k++
	}
	return head
}

// @lc code=end
