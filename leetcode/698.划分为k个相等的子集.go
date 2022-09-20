package leetcode

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=698 lang=golang
 *
 * [698] 划分为k个相等的子集
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) < k {
		return false
	} else if len(nums) == k {
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[0] {
				return false
			}
		}
	}
	sort.Ints(nums)
	fmt.Println(nums)
	avg := 0
	for i := 0; i < len(nums); i++ {
		avg += nums[i]
	}
	if avg%k != 0 {
		return false
	}
	avg /= k
	if nums[len(nums)-1] > avg {
		return false
	}
	count := 0
	checked := map[int]bool{len(nums) - 1: true} //索引：选择了
	var backtrace func(total int)
	backtrace = func(total int) {
		if count == k {
			return
		}
		if total == avg {
			count++
			if count == k {
				return
			}
			for i := len(nums) - 2; i >= 0; i-- {
				if !checked[i] {
					backtrace(nums[i])
					return
				}
			}
			return
		}
		for i := 0; i < len(nums)-1; i++ {
			if !checked[i] {
				backtrace(total + nums[i])
			}
		}
	}
	backtrace(nums[len(nums)-1])
	fmt.Println(count)
	return count == k
}

// @lc code=end
