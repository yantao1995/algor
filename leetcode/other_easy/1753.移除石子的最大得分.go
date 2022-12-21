package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1753 lang=golang
 *
 * [1753] 移除石子的最大得分
 */

// @lc code=start
func maximumScore(a int, b int, c int) int {
	result := 0
	count := 0
	temp := []int{a, b, c}
	for {
		count = 0
		for k := range temp {
			if temp[k] == 0 {
				count++
			}
		}
		if count > 1 {
			break
		}
		sort.Ints(temp)
		temp[1]--
		temp[2]--
		result++
	}
	return result
}

// @lc code=end

/*
	直接模拟整个过程即可
*/
