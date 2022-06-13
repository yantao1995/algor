package leetcode

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * } */

//
func reverseKGroup(head *ListNode, k int) *ListNode {
	oldHead := head
	list := []*ListNode{}
	i := 0
	for head != nil && i < k {
		i++
		list = append(list, head)
		head = head.Next
	}
	if i < k || i == 0 {
		return head
	}
	head3 := list[len(list)-1]
	prev := list[0]
	head2 := oldHead
	temp := 0
	first := true
	for head2 != nil {
		list[temp] = head2
		temp++
		if temp == len(list) { //开始翻转
			list[0].Next = list[len(list)-1].Next
			for m := len(list) - 1; m > 0; m-- {
				if list[m] != nil {
					list[m].Next = list[m-1]
				}
			}
			temp = 0
			if !first {
				prev.Next = list[len(list)-1]
			} else {
				first = false
			}
			prev = list[0]
			head2 = prev
			for n := 0; n < len(list); n++ {
				list[n] = nil
			}
		}
		head2 = head2.Next
	}
	return head3
}

// @lc code=end
