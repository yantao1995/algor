package leetcode

/*
 * @lc app=leetcode.cn id=1790 lang=golang
 *
 * [1790] 仅执行一次字符串交换能否使两个字符串相等
 */

// @lc code=start
func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	count := 0
	i1, i2 := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if count == 2 {
				return false
			}
			if count == 0 {
				i1 = i
			} else {
				i2 = i
			}
			count++
		}
	}
	if count == 2 {
		return s1[i2] == s2[i1] && s1[i1] == s2[i2]
	}
	return count == 0
}

// @lc code=end

/*
	记录count，仅能交换一次，count 为0或2
	i1,i2 记录两次不一样的位置
	然后 判断 相互交换 能否相等
*/
