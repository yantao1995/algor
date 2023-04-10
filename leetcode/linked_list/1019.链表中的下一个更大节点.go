package leetcode

/*
 * @lc app=leetcode.cn id=1019 lang=golang
 *
 * [1019] 链表中的下一个更大节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nextLargerNodes(head *ListNode) []int {
	result := []int{}
	for ; head != nil; head = head.Next {
		result = append(result, 0)
		for hn := head.Next; hn != nil; hn = hn.Next {
			if hn.Val > head.Val {
				result[len(result)-1] = hn.Val
				break
			}
		}
	}
	return result
}

// @lc code=end

/*
	双重循环遍历即可
*/
