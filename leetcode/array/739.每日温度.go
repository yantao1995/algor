package leetcode

/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for k := range temperatures {
		for j := k + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[k] {
				result[k] = j - k
				break
			}
		}
	}
	return result
}

// @lc code=end

/*
	双重循环，来找比i天高的气温天j
*/
