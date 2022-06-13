package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 */

// @lc code=start

func longestWord(words []string) string {
	ht := map[string]int{}
	sort.Strings(words)
	maxStr, length := "", 0
	for _, v := range words {
		if _, ok := ht[string(v[:len(v)-1])]; ok {
			ht[v] = ht[string(v[:len(v)-1])] + 1
			if ht[v] == len(v) && length < len(v) {
				maxStr = v
				length = len(v)
			}
		} else {
			ht[v] = 1
			if len(v) == 1 && length < len(v) {
				maxStr = v
				length = len(v)
			}
		}
	}
	return maxStr
}

// @lc code=end
