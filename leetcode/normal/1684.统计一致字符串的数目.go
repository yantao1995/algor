package leetcode

/*
 * @lc app=leetcode.cn id=1684 lang=golang
 *
 * [1684] 统计一致字符串的数目
 */

// @lc code=start
func countConsistentStrings(allowed string, words []string) int {
	m := map[byte]bool{}
	for i := 0; i < len(allowed); i++ {
		m[allowed[i]] = true
	}
	result := 0
lab:
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if !m[words[i][j]] {
				continue lab
			}
		}
		result++
	}
	return result
}

// @lc code=end

/*
	直接模拟整个过程即可
*/
