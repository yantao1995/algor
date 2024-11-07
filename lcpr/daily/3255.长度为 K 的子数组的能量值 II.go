package lcpr

/*
 * @lc app=leetcode.cn id=3255 lang=golang
 * @lcpr version=30204
 *
 * [3255] 长度为 K 的子数组的能量值 II
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func resultsArray(nums []int, k int) []int {
	result := make([]int, len(nums)-k+1)
	count := make([]int, len(nums)) //记录后一个比前一个大的累加值
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] == 1 {
			count[i] = count[i-1] + 1
		}
	}
	for i := k - 1; i < len(count); i++ {
		result[i-k+1] = -1
		if count[i] >= k-1 { //3个数的话，那值应该是 相差2
			result[i-k+1] = nums[i]
		}
	}
	return result
}

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

/*
	相比于难度1，优化时间复杂度，用count数组记录当前的连续值
	然后遍历的时候，判断当前连续值是否大于等于k-1即可。
	>=k-1 是因为  [0,1,2,3] ,如果3个数，相差值为2就有3个数连续
*/
