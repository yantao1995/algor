package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=491 lang=golang
 *
 * [491] 递增子序列
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	result := [][]int{}
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
			if len(arr) > 1 {
				result = append(result, append([]int{}, arr...))
			}
			distinct[str] = true
			for ; i < len(nums); i++ {
				if len(arr) == 0 || nums[i] >= arr[len(arr)-1] {
					backtrace(i+1, append(arr, nums[i]))
				}
			}
		}
	}
	backtrace(0, []int{})
	return result
}

// @lc code=end

/*
	子集-ii  的剪枝版本
	剪枝： 当前元素如果小于 arr 的末尾元素，就剪枝
*/
