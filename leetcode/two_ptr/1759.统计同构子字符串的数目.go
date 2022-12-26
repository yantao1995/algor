package leetcode

/*
 * @lc app=leetcode.cn id=1759 lang=golang
 *
 * [1759] 统计同构子字符串的数目
 */

// @lc code=start
func countHomogenous(s string) int {
	mod := int(1e9 + 7)
	result := 0
	for i := 0; i < len(s); {
		length := 1
		for length+i < len(s) && s[i+length] == s[i] {
			length++
		}
		i += length
		for length > 0 {
			result += length
			length--
			result %= mod
		}
	}
	return result
}

// @lc code=end

/*
	一段连续长度的子字符串数目计算公式
	例如：zzzzz  = 5+4+3+2+1
	那么只需要双指针找到当前最大的连续长度
		即可算出当前相同的子字符串的数目
	然后将i跳转到下一个不同的字符串位置
*/
