package leetcode

/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
func nthUglyNumber(n int) int {
	dp := make([]int, 1, n)
	dp[0] = 1
	m := map[int]bool{1: true}
	p2, p3, p5 := 0, 0, 0
	v2, v3, v5 := 2, 3, 5
	for i := 0; i < n; {
		if v2 <= v3 && v2 <= v5 {
			if !m[v2] {
				dp = append(dp, v2)
				m[v2] = true
				i++
			}
			p2++
			v2 = dp[p2] * 2

		} else if v3 <= v2 && v3 <= v5 {
			if !m[v3] {
				dp = append(dp, v3)
				m[v3] = true
				i++
			}
			p3++
			v3 = dp[p3] * 3
		} else if v5 <= v2 && v5 <= v3 {
			if !m[v5] {
				dp = append(dp, v5)
				m[v5] = true
				i++
			}
			p5++
			v5 = dp[p5] * 5
		}

	}
	return dp[n-1]
}

// @lc code=end

/*
	如果直接每次将每个数字依次 *2 *3 *5 ，会超时
	所以记录 2 3 5 的乘积指针和 当前乘积的值
	然后保持dp数组有序，依次添加当前指针的最小值，
	然后m用来去重
*/
