package leetcode

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //
func deleteDuplicates(head *ListNode) *ListNode {
	ht := map[int]int{}
	for h := head; h != nil; h = h.Next {
		ht[h.Val]++
	}
	temph := &ListNode{}
	temph.Next = head
	for h := temph; h != nil; {
		if h.Next != nil && ht[h.Next.Val] > 1 {
			h.Next = h.Next.Next
			continue
		}
		h = h.Next
	}
	return temph.Next
}

// @lc code=end
