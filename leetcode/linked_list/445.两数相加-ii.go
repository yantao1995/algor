package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=445 lang=golang
 *
 * [445] 两数相加 II
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
	num1, num2 := "", ""
	for h1 := l1; h1 != nil; h1 = h1.Next {
		num1 += strconv.Itoa(h1.Val)
	}
	for h2 := l2; h2 != nil; h2 = h2.Next {
		num2 += strconv.Itoa(h2.Val)
	}
	result := strAdd445([]byte(num1), []byte(num2))
	num := strings.Split(string(result), "")
	head := &ListNode{}
	h := head
	for i := 0; i < len(num); i++ {
		iT, _ := strconv.Atoi(num[i])
		newNode := &ListNode{
			Val:  iT,
			Next: nil,
		}
		head.Next = newNode
		head = head.Next
	}
	return h.Next
}
func strAdd445(a, b []byte) []byte {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	zero := []byte{'0'}
	result := []byte{}
	weight := zero[0] - 48
	for index := 1; index <= len(a) || index <= len(b); index++ {
		var temp byte
		if index <= len(a) {
			temp += (a[len(a)-index] - 48)
		}
		if index <= len(b) {
			temp += (b[len(b)-index] - 48)
		}
		temp += weight
		if temp >= 10 {
			weight = 1
		} else {
			weight = 0
		}
		bit := []byte{temp%10 + 48}
		result = append(bit, result...)
	}
	if weight > 0 {
		result = append([]byte{'1'}, result...)
	}
	for k := 0; k < len(result); k++ { //去首位0
		if result[k] == '0' {
			result = result[k+1:]
			k--
		} else {
			return result
		}
	}
	if len(result) == 0 {
		return []byte{'0'}
	}
	return result
}

// @lc code=end
