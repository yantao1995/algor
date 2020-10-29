package leetcode

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	if len(s) < 2 {
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
	maxLeft, maxRight, maxLength := 0, 0, 1
	for i := 1; i < len(s); i++ { //长度2开始
		for j := 0; j+i < len(s); j++ {
			if dist[j+1][j+i-1] || i == 1 { //字串是
				if s[j] == s[j+i] {
					dist[j][j+i] = true
					if i+1 > maxLength {
						maxLeft, maxRight, maxLength = j, j+i, i
					}
				}
			}
		}
	}
	return string(s[maxLeft : maxRight+1])
}

// @lc code=end
