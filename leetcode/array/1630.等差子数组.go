package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=1630 lang=golang
 *
 * [1630] 等差子数组
 */

// @lc code=start
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	result := make([]bool, len(l))
	temp := make([]int, len(nums))
next:
	for k := range l {
		temp = append([]int(nil), nums[l[k]:r[k]+1]...)
		sort.Ints(temp)
		t := temp[1] - temp[0]
		for j := 2; j < len(temp); j++ {
			if temp[j]-temp[j-1] != t {
				continue next
			}
		}
		result[k] = true
	}
	return result
}

// @lc code=end

/*
	直接按题目模拟即可
*/
