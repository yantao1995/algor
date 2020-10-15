package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=824 lang=golang
 *
 * [824] 山羊拉丁文
 */

// @lc code=start
func toGoatLatin(S string) string {
	ht := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	slice := strings.Split(S, " ")
	result := ""
	for k := range slice {
		if !ht[slice[k][0]] {
			slice[k] = slice[k][1:] + slice[k][:1]
		}
		result += slice[k] + "ma"
		for i := -1; i < k; i++ {
			result += "a"
		}
		if k != len(slice)-1 {
			result += " "
		}
	}
	return result
}

// @lc code=end
