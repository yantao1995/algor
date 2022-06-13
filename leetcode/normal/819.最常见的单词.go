package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=819 lang=golang
 *
 * [819] 最常见的单词
 */

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
	ht, banHT := map[string]int{}, map[string]bool{}
	paragraph = strings.ToLower(paragraph)
	slice := strings.FieldsFunc(paragraph, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.' || r == '!' || r == '?' || r == ';' || r == '\''
	})
	for k := range banned {
		banHT[banned[k]] = true
	}
	noBanSlice := []string{}
	for k := range slice {
		if !banHT[slice[k]] {
			ht[slice[k]]++
			noBanSlice = append(noBanSlice, slice[k])
		}
	}
	maxVal, result := 0, ""
	for k := range noBanSlice {
		if maxVal < ht[noBanSlice[k]] {
			result = noBanSlice[k]
			maxVal = ht[result]
		}
	}
	return result
}

// @lc code=end
