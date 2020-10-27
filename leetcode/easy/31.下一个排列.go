package easy

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for k := len(nums) - 2; k >= 0; k-- {
		for m := len(nums) - 1; m > k; m-- {
			if nums[m] > nums[k] {
				nums[m], nums[k] = nums[k], nums[m]
				QuickSort31(nums, k+1, len(nums)-1)
				return
			}
		}
	}
	QuickSort31(nums, 0, len(nums)-1)
}
func QuickSort31(nums []int, low, high int) {
	if low >= high {
		return
	}
	i := low + 1
	j := high
	sufftoprev := true
	index := low
	for i <= j {
		if sufftoprev { //坑在前面，哨兵从后往前
			if nums[index] > nums[j] {
				nums[index], nums[j] = nums[j], nums[index]
				index = j
				sufftoprev = false
			}
			j--
		} else { //坑在后面，哨兵从前往后
			if nums[index] < nums[i] {
				nums[index], nums[i] = nums[i], nums[index]
				index = i
				sufftoprev = true
			}
			i++
		}
	}
	QuickSort31(nums, low, index-1)
	QuickSort31(nums, index+1, high)
}

// @lc code=end
