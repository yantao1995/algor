package leetcode

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	nums3 := make([]int, length)
	index := 0
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				nums3[index] = nums1[i]
				i++
			} else {
				nums3[index] = nums2[j]
				j++
			}
		} else if j >= len(nums2) {
			nums3[index] = nums1[i]
			i++
		} else {
			nums3[index] = nums2[j]
			j++
		}

		index++
	}
	if length%2 == 0 {
		return (float64(nums3[length/2-1]) + float64(nums3[length/2])) / 2
	}
	return float64(nums3[length/2])
}

// @lc code=end
