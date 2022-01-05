package leetcode

import (
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode.cn id=382 lang=golang
 *
 * [382] 链表随机节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	ListNode *ListNode
	Len      int
	Rand     *rand.Rand
}

func Constructor(head *ListNode) Solution {
	s := Solution{
		ListNode: head,
		Len:      0,
		Rand:     rand.New(rand.NewSource(time.Now().Unix())),
	}
	for h := head; h != nil; h = h.Next {
		s.Len++
	}
	return s
}

func (this *Solution) GetRandom() int {
	if this.Len > 0 {
		rd := this.Rand.Intn(this.Len) + 1
		for h := this.ListNode; h != nil; h = h.Next {
			rd--
			if rd == 0 {
				return h.Val
			}
		}
	}
	return 0
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end
