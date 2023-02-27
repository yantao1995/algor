package leetcode

/*
 * @lc app=leetcode.cn id=1144 lang=golang
 *
 * [1144] 递减元素使数组呈锯齿状
 */

// @lc code=start
func movesToMakeZigzag(nums []int) int {
	count1, count2 := 0, 0
	nums1 := append([]int{nums[0] - 1}, append(nums, nums[len(nums)-1]-1)...)
	nums2 := make([]int, len(nums1))
	copy(nums2, nums1)
	for i := 1; i < len(nums1)-1; i += 2 {
		if nums1[i-1] >= nums1[i] {
			count1 += nums1[i-1] - nums1[i] + 1
		}
		if nums1[i+1] >= nums1[i] {
			count1 += nums1[i+1] - nums1[i] + 1
			nums1[i+1] = nums1[i] - 1
		}

	}
	for i := 2; i < len(nums2)-1; i += 2 {
		if nums2[i-1] >= nums2[i] {
			count2 += nums2[i-1] - nums2[i] + 1
		}
		if nums2[i+1] >= nums2[i] {
			count2 += nums2[i+1] - nums2[i] + 1
			nums2[i+1] = nums2[i] - 1
		}

	}
	if count1 < count2 {
		return count1
	}
	return count2
}

// @lc code=end

/*
	头和尾添加符合条件的元素，方便处理，然后直接模拟，
	模拟时直接将第 i+1 个元素变成比第 i 个元素小一的值
*/
