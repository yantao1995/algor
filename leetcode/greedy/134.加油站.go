package leetcode

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if gas[i] > 0 {
			last := gas[i]
			j := i
			for last >= cost[j] {
				last -= cost[j]
				if j+1 == len(gas) {
					j = 0
				} else {
					j++
				}
				if i == j {
					return i
				}
				last += gas[j]
			}
		}
	}
	return -1
}

// @lc code=end

/*
	只要当前 得到 gis[j] + 剩余last >= cost[j] 就可以继续往下走
*/
