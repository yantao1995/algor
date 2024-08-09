package lcpr

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=3132 lang=golang
 * @lcpr version=30204
 *
 * [3132] 找出与数组相加的整数 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func minimumAddedInteger(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	sum1, sum2 := 0, 0
	for i := 0; i < len(nums1); i++ {
		sum1 += nums1[i]
		if i < len(nums2) {
			sum2 += nums2[i]
		}
	}
	isSuccess := func(a, b []int) (int, bool) {
		sub := b[0] - a[0]
		for i := 1; i < len(a); i++ {
			if b[i]-a[i] != sub {
				return 0, false
			}
		}
		return sub, true
	}
	result := math.MaxInt
	temp := make([]int, len(nums2))
	for i := 0; i < len(nums1); i++ {
		for j := i + 1; j < len(nums1); j++ {
			temp = temp[:0]
			for m := range nums1 {
				if m != i && m != j {
					temp = append(temp, nums1[m])
				}
			}
			if sub, ok := isSuccess(temp, nums2); ok {
				result = min(result, sub)
			}

		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [4,20,16,12,8]\n[14,18,10]\n
// @lcpr case=end

// @lcpr case=start
// [3,5,5,3]\n[7,7]\n
// @lcpr case=end

*/

/*
	排序后列举出所有情况，然后判断即可
*/
