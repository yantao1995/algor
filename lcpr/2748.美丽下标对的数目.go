package lcpr

/*
 * @lc app=leetcode.cn id=2748 lang=golang
 * @lcpr version=30204
 *
 * [2748] 美丽下标对的数目
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func countBeautifulPairs(nums []int) int {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	tail := make([]int, len(nums))
	for k := range nums {
		tail[k] = nums[k] % 10
		for nums[k] > 10 {
			nums[k] /= 10
		}
	}
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == 1 || tail[j] == 1 || gcd(nums[i], tail[j]) == 1 {
				result++
			}
		}
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [2,5,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [11,21,12]\n
// @lcpr case=end

*/
