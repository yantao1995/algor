package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start

func permuteUnique(nums []int) [][]int {
	total := [][]int{}
	temp := []int{}
	distinctMap := map[string]bool{}
	backTrackUnique(nums, temp, distinctMap, &total)
	return total
}
func backTrackUnique(nums, temp []int, distinctMap map[string]bool, total *[][]int) {
	if len(temp) == len(nums) {
		result := ""
		for k := range temp {
			result += strconv.Itoa(nums[temp[k]]) + ","
		}
		if !distinctMap[result] {
			distinctMap[result] = true
			temp2 := make([]int, len(nums))
			for k := range temp2 {
				temp2[k] = nums[temp[k]]
			}
			*total = append(*total, temp2)
		}
		return
	}
	for k := 0; k < len(nums); k++ {
		if isExist47(temp, k) {
			continue
		}
		temp = append(temp, k)
		backTrackUnique(nums, temp, distinctMap, total)
		temp = temp[:len(temp)-1]
	}
}
func isExist47(temp []int, index int) bool {
	for k := range temp {
		if temp[k] == index {
			return true
		}
	}
	return false
}

// @lc code=end
