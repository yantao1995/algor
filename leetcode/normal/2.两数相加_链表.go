package leetcode

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */ //
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	isTen := 0 //上一轮是否大于10
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			temp := l1.Val + l2.Val + isTen
			p.Val = temp % 10
			isTen = 0
			if temp >= 10 {
				isTen = 1
			}
			l1, l2 = l1.Next, l2.Next
			if l1 != nil || l2 != nil {
				p.Next = &ListNode{}
				p = p.Next
			}
			continue
		}
		if l2 == nil { //只计算l1
			temp := l1.Val + isTen
			p.Val = temp % 10
			isTen = 0
			if temp >= 10 {
				isTen = 1
			}
			l1 = l1.Next
			if l1 != nil {
				p.Next = &ListNode{}
				p = p.Next
			}
			continue
		}
		if l1 == nil { //只计算l1
			temp := l2.Val + isTen
			p.Val = temp % 10
			isTen = 0
			if temp >= 10 {
				isTen = 1
			}
			l2 = l2.Next
			if l2 != nil {
				p.Next = &ListNode{}
				p = p.Next
			}
			continue
		}
	}
	if isTen == 1 {
		p.Next = &ListNode{
			Val:  isTen,
			Next: nil,
		}

	}
	return head
}

// @lc code=end
