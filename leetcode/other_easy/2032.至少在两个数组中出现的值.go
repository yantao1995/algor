package leetcode

/*
 * @lc app=leetcode.cn id=2032 lang=golang
 *
 * [2032] 至少在两个数组中出现的值
 */

// @lc code=start
func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	result := []int{}
	m, m2, m3 := map[int]int{}, map[int]bool{}, map[int]bool{}
	for _, v := range nums1 {
		if m[v] == 0 {
			m[v]++
		}
	}
	for _, v := range nums2 {
		m2[v] = true
	}
	for _, v := range nums3 {
		m3[v] = true
	}
	for k := range m2 {
		if m[k] == 1 {
			result = append(result, k)
		}
		m[k]++
	}
	for k := range m3 {
		if m[k] == 1 {
			result = append(result, k)
		}
		m[k]++
	}
	return result
}

// @lc code=end

/*
	存map，然后统计次数即可
*/
