package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1691 lang=golang
 *
 * [1691] 堆叠长方体的最大高度
 */

// @lc code=start
func maxHeight(cuboids [][]int) int {
	for i := 0; i < len(cuboids); i++ {
		sort.Ints(cuboids[i])
	}
	sort.Slice(cuboids, func(i, j int) bool {
		return cuboids[i][0] < cuboids[j][0] ||
			(cuboids[i][0] == cuboids[j][0] && cuboids[i][1] < cuboids[j][1]) ||
			(cuboids[i][0] == cuboids[j][0] && cuboids[i][1] == cuboids[j][1] && cuboids[i][2] < cuboids[j][2])

	})
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	result := 0
	dp := make([]int, len(cuboids))
	for i := 0; i < len(cuboids); i++ {
		dp[i] = cuboids[i][2]
		for j := 0; j < i; j++ {
			if cuboids[i][1] >= cuboids[j][1] && cuboids[i][2] >= cuboids[j][2] {
				dp[i] = max(dp[i], dp[j]+cuboids[i][2])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

// @lc code=end

/*
	3元组升序排列，然后cuboids按照第一个，第二个，第三个依次升序排序、
	dp[i] 到以第i个结尾的最大高度，所以可能 dp[i]并不是最高的。所以需要记录当前的最大值
*/
