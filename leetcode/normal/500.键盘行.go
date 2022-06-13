package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=500 lang=golang
 *
 * [500] 键盘行
 */

// @lc code=start
func findWords(words []string) []string {
	ht := map[byte]int{
		'z': 1,
		'x': 1,
		'c': 1,
		'v': 1,
		'b': 1,
		'n': 1,
		'm': 1,
		'a': 2,
		's': 2,
		'd': 2,
		'f': 2,
		'g': 2,
		'h': 2,
		'j': 2,
		'k': 2,
		'l': 2,
	}
	result := []string{}
	for k := range words {
		tempWord := strings.ToLower(words[k])
		for b := range tempWord {
			tempInt := ht[tempWord[0]]
			if ht[tempWord[b]] != tempInt {
				goto come
			}
		}
		result = append(result, words[k])
	come:
	}
	return result
}

// @lc code=end
