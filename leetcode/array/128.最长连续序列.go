package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	max := 0
	m := map[int]struct{}{}
	for k := range nums {
		m[nums[k]] = struct{}{}
	}
	nums2 := make([]int, len(m))
	nums2 = nums2[0:0]
	for k := range m {
		nums2 = append(nums2, k)
	}
	nums = nums2
	sort.Ints(nums)
	for i := 0; i < len(nums)-max; i++ {
		if max < 1 {
			max = 1
		}
		current := 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[j-1] == 1 {
				current++
				if current > max {
					max = current
				}
				continue
			}
			i = j - 1
			break
		}
	}
	return max
}

// @lc code=end

/*
	先对数组排序，然后依次找最长序列。遇到不连续，i指针直接跳过来。当前序列如果超过未遍历值，也跳出。
*/
