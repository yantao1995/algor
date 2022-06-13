package leetcode

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	for n2Index := 0; n2Index < n; n2Index++ {
		for n1Index := 0; n1Index <= m; n1Index++ {
			if nums1[n1Index] > nums2[n2Index] {
				for i := m - 1; i >= n1Index; i-- {
					nums1[i+1] = nums1[i]
				}
				nums1[n1Index] = nums2[n2Index]
				m++
				break
			}
			if n1Index == m {
				nums1[n1Index] = nums2[n2Index]
				m++
				break
			}
		}
	}
}

// @lc code=end
