package leetcode

/*
 * @lc app=leetcode.cn id=1775 lang=golang
 *
 * [1775] 通过最少操作次数使数组的和相等
 */

// @lc code=start
/* func minOperations(nums1 []int, nums2 []int) int {
	sum1, sum2 := 0, 0
	for k := range nums1 {
		sum1 += nums1[k]
	}
	for k := range nums2 {
		sum2 += nums2[k]
	}
	if sum1 == sum2 {
		return 0
	} else if sum1 > sum2 {
		nums1, nums2 = nums2, nums1
		sum1, sum2 = sum2, sum1
	}
	sort.Ints(nums1)
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] > nums2[j]
	})
	result := 0
	tempI, tempJ := -1, -1
	for i, j := 0, 0; i < len(nums1) || j < len(nums2); tempI, tempJ = -1, -1 {
		if i < len(nums1) {
			tempI = 6 - nums1[i]
		}
		if j < len(nums2) {
			tempJ = nums2[j] - 1
		}
		result++
		if tempI >= tempJ {
			sum1 += tempI
			i++
		} else {
			sum2 -= tempJ
			j++
		}
		if sum1 >= sum2 {
			return result
		}
	}
	return -1
} */

// @lc code=end

/*
	先求sum1,sum2，然后保证 sum1小于等于sum2
	然后 nums1 升序， nums2 降序
	按顺序 nums1[i] 变成6， nums2[j] 变成1，得到变化的差值，每次都取最大的
	看看是否达成 sum1>=sum2 ，如果达成，那说明可以完成操作
*/
