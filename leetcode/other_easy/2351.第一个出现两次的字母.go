package leetcode

/*
 * @lc app=leetcode.cn id=2351 lang=golang
 *
 * [2351] 第一个出现两次的字母
 */

// @lc code=start
func repeatedCharacter(s string) byte {
	log := make([]bool, 26)
	idx := 0
	for i := 0; i < len(s); i++ {
		idx = int(s[i] - 97)
		if log[idx] {
			return s[i]
		}
		log[idx] = true
	}
	return 'a'
}

// @lc code=end

/*
	log长度26，分别记录26个字母是否出现
	idx为字母与log的索引映射
	只要当前字母出现过就直接return
*/
