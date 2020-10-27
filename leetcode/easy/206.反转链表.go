package easy

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //

func reverseList(head *ListNode) *ListNode { //直接逆序
	var prev, next *ListNode = nil, nil
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// func reverseList(head *ListNode) *ListNode { //递归
// 	if head == nil {
// 		return head
// 	}
// 	if head.Next != nil {
// 		newLink := reverseList(head.Next)
// 		head.Next.Next = head
// 		head.Next = nil
// 		return newLink
// 	}
// 	return head
// }

// @lc code=end
