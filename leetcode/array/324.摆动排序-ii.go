package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=324 lang=golang
 *
 * [324] 摆动排序 II
 */

// @lc code=start
func wiggleSort(nums []int) {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	sort.Ints(numsCopy)
	l, r := (len(nums)+1)/2-1, len(nums)-1
	for i := 0; i < len(numsCopy); i++ {
		if i&1 == 0 {
			nums[i] = numsCopy[l]
			l--
		} else {
			nums[i] = numsCopy[r]
			r--
		}
	}
}

// @lc code=end

/*
	因为题目的数据是满足要求的，所以可以对排序后的数分两块，
	然后对两块数据同时进行降序交替排列

	例如  [1,5,5,6] ,如果升序排列， 1  5  5 6 ，则不满足
					降序：  5,6,1,5
*/
