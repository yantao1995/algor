package lcpr

/*
 * @lc app=leetcode.cn id=3038 lang=golang
 * @lcpr version=30203
 *
 * [3038] 相同分数的最大操作数目 I
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxOperations(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	result := 1
	for i := 2; i < len(nums)-1; i += 2 {
		if nums[i]+nums[i+1] != nums[i-1]+nums[i-2] {
			break
		}
		result++
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,1,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,6,1,4]\n
// @lcpr case=end

*/

/*
直接按规则数就可以了
*/
