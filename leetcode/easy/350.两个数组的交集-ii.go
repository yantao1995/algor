package easy

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	mp2 := map[int]int{}
	for k := range nums1 {
		if _, ok := mp[nums1[k]]; ok {
			mp[nums1[k]]++
		} else {
			mp[nums1[k]] = 1
		}
	}
	for k := range nums2 {
		if _, ok := mp[nums2[k]]; ok {
			if mp2[nums2[k]] < mp[nums2[k]] {
				mp2[nums2[k]]++
			}
		}
	}
	temp := []int{}
	for k := range mp2 {
		for mp2[k] > 0 {
			temp = append(temp, k)
			mp2[k]--
		}
	}
	return temp
}

// @lc code=end
