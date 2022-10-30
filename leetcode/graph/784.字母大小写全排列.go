package leetcode

/*
 * @lc app=leetcode.cn id=784 lang=golang
 *
 * [784] 字母大小写全排列
 */

// @lc code=start
func letterCasePermutation(s string) []string {
	result := []string{}
	var distant byte = 32
	var dfs func(bts []byte, i int)
	dfs = func(bts []byte, i int) {
		result = append(result, string(bts))
		if i < len(bts) {
			for ; i < len(bts); i++ {
				if bts[i] >= 'a' {
					bts[i] -= distant
					dfs(bts, i+1)
					bts[i] += distant
				} else if bts[i] >= 'A' {
					bts[i] += distant
					dfs(bts, i+1)
					bts[i] -= distant
				}
			}
		}
	}
	dfs([]byte(s), 0)
	return result
}

// @lc code=end

/*
	dfs 直接枚举所有结果即可
	遇到 >= 'a' 的直接减去32
	遇到 >= 'A' 的直接加上32

	优先判断'a',其次判断'A'
	因为:
			'a' 的ascii码为 97
			'A' 的ascii码为 65
			'0' 的ascii码为 30

*/
