package easy

import "strings"

/*
 * @lc app=leetcode.cn id=1455 lang=golang
 *
 * [1455] 检查单词是否为句中其他单词的前缀
 */

// @lc code=start
func isPrefixOfWord(sentence string, searchWord string) int {
	slice := strings.Split(sentence, " ")
	for k := range slice {
		if len(slice[k]) >= len(searchWord) &&
			slice[k][:len(searchWord)] == searchWord {
			return k + 1
		}
	}
	return -1
}

// @lc code=end
