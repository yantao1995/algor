package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1798 lang=golang
 *
 * [1798] 你能构造出连续值的最大数目
 */

// @lc code=start
func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	result := 0
	for i := 0; i < len(coins); i++ {
		if coins[i] <= result+1 {
			result += coins[i]
			continue
		}
		break
	}
	return result + 1
}

// @lc code=end

/*
	贪心，连续的数是从0开始，result表示当前可达的最大值
	如果加上当前数可以到达下一个数，那么更新result
*/
