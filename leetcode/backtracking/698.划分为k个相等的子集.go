package leetcode

import (
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
	can := false
	checked := make([]bool, len(nums))
	var backtrace func(idx, total, count int)
	backtrace = func(idx, total, count int) {
		if can || total+nums[idx] > avg {
			return
		}
		total += nums[idx]
		if total == avg {
			count++
			if count == k {
				can = true
				return
			}
			total = 0
			idx = len(nums)
		}
		for i := idx - 1; i >= 0; i-- {
			if !checked[i] {
				checked[i] = true
				backtrace(i, total, count)
				checked[i] = false
			}
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		checked[i] = true
		backtrace(i, 0, 0)
		checked[i] = false
	}

	return can
}

// @lc code=end

/*
	回溯爆搜，在回溯前，先排除部分不能满足的答案
	回溯剪枝：
		1.从大到小来加，比从小到大会少搜索很多次，因为比如遇到 [1,1,1,1,1,1,4]这种，
			从后向前就只需要一次，而从前往后，就会一直在1这个上面搜索很多次
		2.can 记录是否搜到了，如果搜到了，直接终止回溯栈
		3.checked 标志数组记录当前索引的数是否已经被添加到了对应组合
		4.如果 total == avg ，那么应该重置 idx = len(nums) ，因为当前idx索引后面还有未加入对应组合的数，
			新组合应该以最大的数为初始值
*/
