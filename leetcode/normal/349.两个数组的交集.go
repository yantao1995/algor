package leetcode

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	mp := map[int]int{}
	for k := range nums1 {
		mp[nums1[k]] = 1
	}
	for k := range nums2 {
		if _, ok := mp[nums2[k]]; ok {
			mp[nums2[k]] = 2
		}
	}
	temp := []int{}
	for k := range mp {
		if mp[k] == 2 {
			temp = append(temp, k)
		}
	}
	return temp
}

// @lc code=end
