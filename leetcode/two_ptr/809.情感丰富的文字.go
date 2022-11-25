package leetcode

/*
 * @lc app=leetcode.cn id=809 lang=golang
 *
 * [809] 情感丰富的文字
 */

// @lc code=start
func expressiveWords(s string, words []string) int {
	result := 0
	canStretchy := func(str string) bool {
		if len(str) > len(s) {
			return false
		}
		i, j := 0, 0
		for i < len(s) && j < len(str) {
			if s[i] != str[j] {
				return false
			}
			a, b := i+1, j+1
			for ; a < len(s) && s[a] == s[i]; a++ {
			}
			for ; b < len(str) && str[b] == str[j]; b++ {
			}
			if a-i != b-j && (a-i < 3 || a-i < b-j) {
				return false
			}
			i, j = a, b
		}
		return i == len(s) && j == len(str)
	}
	for _, str := range words {
		if canStretchy(str) {
			result++
		}
	}
	return result
}

// @lc code=end

/*
	双指针，i指向s,j指向当前str
	如果当前字符相同，那么i和j同时向后移动，一直移动到下一个不同的字符
	这时候，
		如果i,j移动的距离相同，那么直接进入下一个字符判断
		如果不同，
			1. s中i的距离如果小于3，那么说明重复长度不够3，应该返回false
			2. s中的i的距离小于 str中j的距离，那么说明 str中的叠字比 s 中还多，也应该返回false
	最后判断 s中的i 和 str中的j 是否都走到了末尾
*/
