package leetcode

/*
 * @lc app=leetcode.cn id=2469 lang=golang
 *
 * [2469] 温度转换
 */

// @lc code=start
func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.8 + 32}
}

// @lc code=end

/*
	直接写
*/
