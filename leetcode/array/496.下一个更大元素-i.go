package leetcode

/*
 * @lc app=leetcode.cn id=496 lang=golang
 *
 * [496] 下一个更大元素 I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{} //val 索引
	for k := range nums2 {
		m[nums2[k]] = k
	}
	result := make([]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		for j := m[nums1[i]]; j < len(nums2); j++ {
			if nums2[j] > nums1[i] {
				result = append(result, nums2[j])
				goto lab
			}
		}
		result = append(result, -1)
	lab:
	}
	return result
}

// @lc code=end

/*
	用m记录 nums2 中值的索引,
	这样遍历nums1时可以直接定位到nums2的开始位置，然后开始遍历
*/
