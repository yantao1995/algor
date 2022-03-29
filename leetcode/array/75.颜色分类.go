package leetcode

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
// func sortColors(nums []int) {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] > nums[j] {
// 				nums[i], nums[j] = nums[j], nums[i]
// 			}
// 		}
// 	}
// }

func sortColors(nums []int) {
	var quickSort func(low, high int)
	quickSort = func(low, high int) {
		if low >= high {
			return
		}
		i, j := low+1, high
		index := low
		directToTail := true
		for i <= j {
			if directToTail {
				if nums[index] > nums[j] {
					nums[index], nums[j] = nums[j], nums[index]
					index = j
					directToTail = false
				}
				j--
			} else {
				if nums[index] < nums[i] {
					nums[index], nums[i] = nums[i], nums[index]
					index = i
					directToTail = true
				}
				i++
			}
		}
		quickSort(low, index-1)
		quickSort(index+1, high)
	}
	quickSort(0, len(nums)-1)
}

// @lc code=end
