package lcpr

/*
 * @lc app=leetcode.cn id=3115 lang=golang
 * @lcpr version=30204
 *
 * [3115] 质数的最大距离
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maximumPrimeDifference(nums []int) int {
	isPrime := func(n int) bool {
		if n == 1 {
			return false
		}
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	min, max := 0, len(nums)-1
	for min < max {
		if isPrime(nums[min]) {
			break
		}
		min++
	}
	for min < max {
		if isPrime(nums[max]) {
			break
		}
		max--
	}
	return max - min
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,9,5,3]\n
// @lcpr case=end

// @lcpr case=start
// [4,8,2,8]\n
// @lcpr case=end

*/
