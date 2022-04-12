package leetcode

/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	m12, m34 := map[int]int{}, map[int]int{}
	for k := range nums1 {
		for kk := range nums2 {
			m12[nums1[k]+nums2[kk]]++
		}
	}
	for k := range nums3 {
		for kk := range nums4 {
			m34[nums3[k]+nums4[kk]]++
		}
	}
	for s1, ct := range m12 {
		count += ct * m34[0-s1]
	}
	return count
}

// @lc code=end

/*
	用map将四个 for{}  降维到 两个for {}
	n1+n2+n3+n4 =0
	可以转化为  n1 +n2 = -n3-n4
	所以相当于两个map的和作比较
*/
