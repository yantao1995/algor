package leetcode

/*
 * @lc app=leetcode.cn id=890 lang=golang
 *
 * [890] 查找和替换模式
 */

// @lc code=start
func findAndReplacePattern(words []string, pattern string) []string {
	result := []string{}
	isMatch := func(pattern, word string) bool {
		if len(pattern) != len(word) {
			return false
		}
		m := map[byte]byte{}
		distinct := map[byte]bool{}
		for k := range pattern {
			if v, ok := m[pattern[k]]; !ok {
				m[pattern[k]] = word[k]
				if distinct[word[k]] {
					return false
				}
				distinct[word[k]] = true
			} else if v != word[k] {
				return false
			}
		}
		return true
	}
	for k := range words {
		if isMatch(pattern, words[k]) {
			result = append(result, words[k])
		}
	}
	return result
}

// @lc code=end

/*
    m 建立 pattern 与 word 单词的字符映射关系
	distinct 防止 m 中的关系出现重复，比如 pattern = "abb"  word = "ccc"
	遍历pattern与word，
		如果当前字符对没有建立过映射，就建立映射关系，
	 		同时判断word的当前单词是否已经有过映射绑定，如果已经绑定过，那么不匹配
		如果有映射关系，就判断是否相同
*/
