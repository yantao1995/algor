package lcpr

/*
 * @lc app=leetcode.cn id=2956 lang=golang
 * @lcpr version=30204
 *
 * [2956] 找到两个数组中的公共元素
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	result := []int{0, 0}
	m1, m2 := map[int]bool{}, map[int]bool{}
	for _, v := range nums1 {
		m1[v] = true
	}
	for _, v := range nums2 {
		m2[v] = true
		if m1[v] {
			result[1]++
		}
	}
	for _, v := range nums1 {
		if m2[v] {
			result[0]++
		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [4,3,2,3,1]\n[2,2,5,2,3,6]\n
// @lcpr case=end

// @lcpr case=start
// [3,4,2,3]\n[1,5]\n
// @lcpr case=end

*/

/*
仔细读题
*/
