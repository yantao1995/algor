package leetcode

/*
 * @lc app=leetcode.cn id=1491 lang=golang
 *
 * [1491] 去掉最低工资和最高工资后的工资平均值
 */

// @lc code=start
func average(salary []int) float64 {
	sum, min, max := 0, 2<<31-1, 0
	for k := range salary {
		sum += salary[k]
		if min > salary[k] {
			min = salary[k]
		}
		if max < salary[k] {
			max = salary[k]
		}
	}
	return float64(sum-min-max) / float64(len(salary)-2)
}

// @lc code=end
