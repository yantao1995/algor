package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=884 lang=golang
 *
 * [884] 两句话中的不常见单词
 */

// @lc code=start
func uncommonFromSentences(A string, B string) []string {
	slice1, slice2 := strings.Split(A, " "), strings.Split(B, " ")
	ht1, ht2 := map[string]int{}, map[string]int{}
	for k := range slice1 {
		ht1[slice1[k]]++
	}
	for k := range slice2 {
		ht1[slice2[k]]++
	}
	result := []string{}
	for k := range ht1 {
		if ht1[k] == 1 {
			if ht2[k] == 0 {
				result = append(result, k)
			}
		}
	}
	for k := range ht2 {
		if ht2[k] == 1 {
			if ht1[k] == 0 {
				result = append(result, k)
			}
		}
	}
	return result
}

// @lc code=end
