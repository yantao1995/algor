package leetcode

/*
 * @lc app=leetcode.cn id=1704 lang=golang
 *
 * [1704] 判断字符串的两半是否相似
 */

// @lc code=start
func halvesAreAlike(s string) bool {
	count := 0
	m := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
	for i := 0; i < len(s); i++ {
		if m[s[i]] {
			if i < len(s)/2 {
				count++
			} else {
				count--
			}
		}
	}
	return count == 0
}

// @lc code=end

/*
	直接计数即可
*/
