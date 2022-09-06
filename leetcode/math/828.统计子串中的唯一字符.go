package leetcode

/*
 * @lc app=leetcode.cn id=828 lang=golang
 *
 * [828] 统计子串中的唯一字符
 */

// @lc code=start
func uniqueLetterString(s string) int {
	result := 0
	btsIdxs := map[byte][]int{}
	for i := 0; i < len(s); i++ {
		btsIdxs[s[i]] = append(btsIdxs[s[i]], i)
	}
	for _, idx := range btsIdxs {
		idx = append([]int{-1}, append(idx, len(s))...)
		for i := 1; i < len(idx)-1; i++ {
			result += (idx[i] - idx[i-1]) * (idx[i+1] - idx[i])
		}
	}
	return result
}

// @lc code=end

/*
	参考官方题解，计算每个字符的贡献，记录每个字符的位置，
	字符的个数，在于前一个相同的字符位置和下一个相同字符的位置



超时 ：暴力解法，依次枚举每种组合结果
func uniqueLetterString(s string) int {
	result := 0
	count := map[byte]int{}
	repeats := 0
	for i := 0; i < len(s); i++ {
		count = map[byte]int{}
		repeats = 0
		for j := i; j < len(s); j++ {
			count[s[j]]++
			if count[s[j]] > 1 {
				repeats++
				if count[s[j]] == 2 {
					repeats++
				}
			}
			result += j - i + 1 - repeats
		}
	}
	return result
}
*/
