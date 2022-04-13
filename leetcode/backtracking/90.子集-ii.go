package leetcode

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=90 lang=golang
 *
 * [90] 子集 II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	distinct := map[string]bool{}
	arrToKey := func(arr []int) string {
		str := " "
		for k := range arr {
			str += strconv.Itoa(arr[k]) + ","
		}
		return str
	}
	var backtrace func(i int, arr []int)
	backtrace = func(i int, arr []int) {
		if str := arrToKey(arr); !distinct[str] {
			result = append(result, append([]int{}, arr...))
			distinct[str] = true
			for ; i < len(nums); i++ {
				backtrace(i+1, append(arr, nums[i]))
			}
		}
	}
	backtrace(0, []int{})
	return result
}

// @lc code=end

/*
	和子集一样回溯处理
	使用了map来去重
*/
