package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=1078 lang=golang
 *
 * [1078] Bigram 分词
 */

// @lc code=start
func findOcurrences(text string, first string, second string) []string {
	slice := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.' || r == '!' || r == '?' || r == ';' || r == '\''
	})
	result := []string{}
	for i := 1; i < len(slice)-1; i++ {
		if second == slice[i] && first == slice[i-1] {
			result = append(result, slice[i+1])
		}
	}
	return result
}

// @lc code=end
