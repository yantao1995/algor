package leetcode

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	rtn := []int{-1, -1}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (high + low) / 2
		if nums[mid] == target {
			rtn[0], rtn[1] = mid, mid
			prevFlag, nextFlag := true, true
			index := 1
			for {
				if prevFlag {
					if mid-index >= 0 {
						if nums[mid-index] == target {
							rtn[0] = mid - index
						} else {
							prevFlag = false
						}
					} else {
						prevFlag = false
					}
				}
				if nextFlag {
					if mid+index < len(nums) {
						if nums[mid+index] == target {
							rtn[1] = mid + index
						} else {
							nextFlag = false
						}
					} else {
						nextFlag = false
					}
				}
				if !prevFlag && !nextFlag {
					return rtn
				}
				index++
			}
		}
		if nums[high] >= nums[low] { //本区间单调递增
			if nums[mid] > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return rtn
}

// @lc code=end
