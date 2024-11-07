package lcpr

/*
 * @lc app=leetcode.cn id=3254 lang=golang
 * @lcpr version=30204
 *
 * [3254] 长度为 K 的子数组的能量值 I
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
// func resultsArray(nums []int, k int) []int {
// 	result := make([]int, len(nums)-k+1)
// lab:
// 	for i := 0; i+k <= len(nums); i++ {
// 		result[i] = -1
// 		for j := i + 1; j < i+k; j++ {
// 			if nums[j] != nums[j-1]+1 {
// 				continue lab
// 			}
// 		}
// 		result[i] = nums[i+k-1]
// 	}
// 	return result
// }

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,3,2,5]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,2,2,2,2]\n4\n
// @lcpr case=end

// @lcpr case=start
// [3,2,3,2,3,2]\n2\n
// @lcpr case=end

*/
