package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=792 lang=golang
 *
 * [792] 匹配子序列的单词数
 */

// @lc code=start
func numMatchingSubseq(s string, words []string) int {
	result := 0
	m := map[byte][]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]] = append(m[s[i]], i)
	}
	lastIndex := -1
lab:
	for _, word := range words {
		lastIndex = -1
	nextBt:
		for i := 0; i < len(word); i++ {
			t := sort.SearchInts(m[word[i]], lastIndex+1)
			if t < len(m[word[i]]) {
				lastIndex = m[word[i]][t]
				continue nextBt
			}
			continue lab
		}
		result++
	}
	return result
}

// @lc code=end

/*

替换超时的内循环为二分查找，即可满足
for _, idx := range m[word[i]] {
	if idx > lastIndex {
		lastIndex = idx
		continue nextBt
	}
}



超时，记录每个索引元素的位置,word下一个索引的字符在原s中索引大于上一个即可。然后连续向下寻找

func numMatchingSubseq(s string, words []string) int {
	result := 0
	m := map[byte][]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]] = append(m[s[i]], i)
	}
		lastIndex := -1
	lab:
		for _, word := range words {
			lastIndex = -1
		nextBt:
			for i := 0; i < len(word); i++ {
				for _, idx := range m[word[i]] {
					if idx > lastIndex {
						lastIndex = idx
						continue nextBt
					}
				}
				continue lab
			}
			result++
		}
	return result
}
*/
