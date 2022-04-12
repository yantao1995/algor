package leetcode

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	m1, m2 := map[byte]int{}, map[byte]int{}
	for k := range s {
		m1[s[k]]++
	}
	for k := range t {
		m2[t[k]]++
	}
	for k := range m1 {
		if m1[k] != m2[k] {
			return false
		}
	}
	return len(m1) == len(m2)
}

// @lc code=end

/*
	统计字符个数，然后依次判断
*/
