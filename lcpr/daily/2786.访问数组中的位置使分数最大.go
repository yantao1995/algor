package lcpr

/*
 * @lc app=leetcode.cn id=2786 lang=golang
 * @lcpr version=30203
 *
 * [2786] 访问数组中的位置使分数最大
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func maxScore(nums []int, x int) int64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxOdd, maxEven := nums[0], nums[0]
	if nums[0]&1 == 0 { //偶数
		maxOdd -= x
	} else {
		maxEven -= x
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]&1 == 0 { //偶数
			maxEven = max(maxEven+nums[i], maxOdd+nums[i]-x)
		} else {
			maxOdd = max(maxOdd+nums[i], maxEven+nums[i]-x)
		}
	}
	return int64(max(maxOdd, maxEven))
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,6,1,9,2]\n5\n
// @lcpr case=end

// @lcpr case=start
// [2,4,6,8]\n3\n
// @lcpr case=end

*/

/*
优化思路：分别记录当前奇偶的最大的一个值，避免内循环向前遍历


思路：动态规划，每次计算到当前能得到的最大值即可。注意初始化当前值。然后记录一个最大值返回即可
超时
Time Limit Exceeded
737/744 cases passed (N/A)

func maxScore(nums []int, x int) int64 {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	result := make([]int, len(nums))
	result[0] = nums[0]
	maxResult := nums[0]
	for i := 1; i < len(nums); i++ {
		result[i] = -x
		for j := i - 1; j >= 0; j-- {
			if nums[i]&1 == nums[j]&1 { //奇偶性相同
				result[i] = max(result[i], result[j]+nums[i])
			} else {
				result[i] = max(result[i], result[j]+nums[i]-x)
			}
		}
		maxResult = max(maxResult, result[i])
	}
	return int64(maxResult)
}

*/
