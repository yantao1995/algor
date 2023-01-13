package leetcode

/*
 * @lc app=leetcode.cn id=2287 lang=golang
 *
 * [2287] 重排字符形成目标字符串
 */

// @lc code=start
func rearrangeCharacters(s string, target string) int {
	m := map[byte]int{}
	result := len(s)
	for _, v := range target {
		m[byte(v)]++
	}
	m2 := map[byte]int{}
	for _, v := range s {
		m2[byte(v)]++
	}
	for k := range m {
		m2[k] /= m[k]
		if m2[k] < result {
			result = m2[k]
		}
	}
	return result
}

// @lc code=end

/*
	因为s可以任意排列，那么可以直接先用m存储每个字符的个数，
	然后用m2存储s中相同字符的个数，如果需要形成目标字符串，
	那么就需要每个字符都有足够的个数，那么可以直接 m[k] / m2[k]
	得到当前字符能组成的个数，然后result记录最小的个数即是
	答案需要的个数
*/
