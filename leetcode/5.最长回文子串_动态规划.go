package leetcode

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 3 {
		return s
	}
	dist := [][]bool{}
	for i := 0; i < len(s); i++ { //备忘录
		dist = append(dist, []bool{})
		for j := 0; j <= len(s); j++ {
			if i != j {
				dist[i] = append(dist[i], []bool{false}...)
			} else {
				dist[i] = append(dist[i], []bool{true}...) //单元素
			}

		}
	}
	for lns := len(s); lns > 1; lns-- { //字串长度
		// for i, j := 0, lns; i+lns < len(s); i, j = i+1, j-1 { //移动游标
		// 	if dist[i+1][j-1] { //是
		// 		if s[i] == s[j] {
		// 			return s[i : j+1]
		// 		}
		// 	}
		// 	for m := i; m < j; m++ {

		// 	}
		// }
	}
	return string(s[0])
}

// @lc code=end
