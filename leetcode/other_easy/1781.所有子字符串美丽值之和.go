package leetcode

/*
 * @lc app=leetcode.cn id=1781 lang=golang
 *
 * [1781] 所有子字符串美丽值之和
 */

// @lc code=start
func beautySum(s string) int {
	result := 0
	m := map[byte]int{}
	max, min := s[0], s[0]
	for i := 0; i < len(s)-2; i++ {
		m = map[byte]int{}
		max, min = s[i], s[i]
		for j := i; j < len(s); j++ {
			m[s[j]]++
			if m[s[j]] > m[max] {
				max = s[j]
			}
			if m[s[j]] == 1 {
				min = s[j]
			} else if s[j] == min {
				for k, v := range m {
					if v < m[min] {
						min = k
					}
				}
			}
			if m[max]-m[min] > 0 {
				result += m[max] - m[min]
			}
		}
	}
	return result
}

// @lc code=end

/*
	m 记录当前子字符串的每个字符的个数
	max,min分别记录当前子字符串内的最大字符和最小字符
	每个字符个数累加后都可以与max比较
	而min的更新，则依赖于当前字符个数是否为1或者当前字符是否为min字符
		如果当前字符是min字符，那么则需要重新找min字符
*/
