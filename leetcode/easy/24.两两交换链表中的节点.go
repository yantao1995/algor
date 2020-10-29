package leetcode

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * } */

//
func swapPairs(head *ListNode) *ListNode {
	prev := head
	if prev == nil {
		return prev
	}
	current := prev.Next
	if current == nil {
		return prev
	}
	next := current.Next
	if next == nil {
		current.Next = prev
		prev.Next = nil
		return current
	}
	currentLength := 0
	hrtn := current
	current.Next = prev
	prev.Next = next
	head = current
	pprve := head.Next
	for head != nil {
		currentLength++
		current = head
		next = head.Next
		if currentLength%2 == 0 && currentLength != 2 {
			pprve.Next = current
			current.Next = prev
			prev.Next = next
			pprve = prev
		}
		prev = head
		head = next
	}
	return hrtn
}

// @lc code=end
