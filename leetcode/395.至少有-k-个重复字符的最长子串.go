package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有 K 个重复字符的最长子串
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	maxLength := 0
	cm := map[byte]int{}
	count := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		cm[s[i]]++
		count[i] = cm[s[i]]
	}
	fmt.Println(count)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	temp := map[byte]int{}
	need := map[byte]bool{}
	got := map[byte]bool{}
	for i := len(count) - 1; i >= 0; i-- {
		temp = map[byte]int{}
		need = map[byte]bool{}
		got = map[byte]bool{}
		for j := i; j >= 0 && (count[i] >= k || need[s[j]] || got[s[j]]); j-- {
			if count[j] < k && !need[s[j]] && !got[s[j]] {
				i = j
				break
			}
			if temp[s[j]] == 0 {
				need[s[j]] = true
				temp[s[j]] = count[j]
			}
			if temp[s[j]]-count[j]+1 >= k {
				got[s[j]] = true
				delete(need, s[j])
			}
			if len(need) == 0 {
				maxLength = max(maxLength, i-j+1)
			}
		}
	}
	return maxLength
}

// @lc code=end

/* 超时
func longestSubstring(s string, k int) int {
	maxLength := 0
	need := map[byte]bool{}
	count := map[byte]int{}
	for i := 0; i < len(s)-maxLength; i++ {
		need = map[byte]bool{}
		count = map[byte]int{}
		for j := i; j < len(s); j++ {
			if count[s[j]] < k {
				need[s[j]] = true
			}
			count[s[j]]++
			if count[s[j]] >= k {
				delete(need, s[j])
			}
			if len(need) == 0 && j-i+1 > maxLength {
				maxLength = j - i + 1
			}
		}
	}
	return maxLength
}
*/
