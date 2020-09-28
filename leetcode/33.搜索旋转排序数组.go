package leetcode

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
// func search(nums []int, target int) int {
// 	low, high := 0, len(nums)-1
// 	for low <= high {
// 		mid := (high + low) / 2
// 		if nums[mid] == target {
// 			return mid
// 		}
// 		if nums[high] >= nums[low] && (nums[low] <= target && nums[high] >= target) { //本区间单调递增
// 			if nums[mid] > target {
// 				high = mid - 1
// 			} else {
// 				low = mid + 1
// 			}
// 		} else {
// 			if (target > nums[mid] && target > nums[low] && nums[mid] > nums[high]) ||
// 				(target > nums[mid] && target < nums[low] && nums[mid] < nums[low]) ||
// 				(target < nums[mid] && target < nums[low] && nums[mid] > nums[low]) ||
// 				(target < nums[mid] && target < nums[low] && nums[high] <= target) {
// 				low = mid + 1
// 			} else {
// 				high = mid - 1
// 			}
// 		}
// 	}
// 	return -1
// }

// @lc code=end
