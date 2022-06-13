package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 */

// @lc code=start
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	factor := []int{}
	end := int(math.Sqrt(float64(num))) + 1
	i := 1
	for i <= end {
		if num%i == 0 {
			if len(factor) > 0 {
				if factor[len(factor)-1] != i {
					factor = append(factor, i)
				}
			} else {
				factor = append(factor, i)
			}
			if i != end && i != 1 {
				factor = append(factor, num/i)
			}
		}
		i++
	}
	result := 0
	for k := range factor {
		result += factor[k]
	}
	if result != num {
		return false
	}
	return true
}

// @lc code=end
