package leetcode

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	total := [][]int{}
	QuickSort47(nums, 0, len(nums)-1)
	//fmt.Println(nums)
	backTrackUnique(nums, 0, &total)
	return total
}
func backTrackUnique(nums []int, i int, total *[][]int) {
	if i == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*total = append(*total, temp)
	}
	for k := i; k < len(nums); k++ {
		if nums[k] == nums[i] && k != i {
			continue
		}
		if k+1 < len(nums) {
			if nums[k] == nums[k+1] && k != i {
				continue
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
		backTrackUnique(nums, i+1, total)
		nums[i], nums[k] = nums[k], nums[i]
	}
}
func QuickSort47(nums []int, low, high int) {
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
	QuickSort47(nums, low, index-1)
	QuickSort47(nums, index+1, high)
}

// @lc code=end
