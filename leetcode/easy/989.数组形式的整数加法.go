package easy

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=989 lang=golang
 *
 * [989] 数组形式的整数加法
 */

// @lc code=start
func addToArrayForm(A []int, K int) []int {
	kSlice := strings.Split(strconv.Itoa(K), "")
	intSlice := make([]int, len(kSlice))
	for k := range kSlice {
		intSlice[k], _ = strconv.Atoi(kSlice[k])
	}
	result := []int{}

	iIndex, jIndex := len(A)-1, len(intSlice)-1
	bit := 0
	for iIndex >= 0 || jIndex >= 0 {
		if iIndex >= 0 && jIndex >= 0 {
			result = append([]int{(A[iIndex] + intSlice[jIndex] + bit) % 10}, result...)
			bit = (A[iIndex] + intSlice[jIndex] + bit) / 10
			iIndex--
			jIndex--
		} else if iIndex < 0 {
			result = append([]int{(intSlice[jIndex] + bit) % 10}, result...)
			bit = (intSlice[jIndex] + bit) / 10
			jIndex--
		} else {
			result = append([]int{(A[iIndex] + bit) % 10}, result...)
			bit = (A[iIndex] + bit) / 10
			iIndex--
		}
	}
	if bit > 0 {
		result = append([]int{1}, result...)
	}
	return result
}

// @lc code=end
