package leetcode

/*
 * @lc app=leetcode.cn id=1658 lang=golang
 *
 * [1658] 将 x 减到 0 的最小操作数
 */

// @lc code=start
func minOperations(nums []int, x int) int {
	n := len(nums)
	if nums[0] == x || nums[n-1] == x {
		return 1
	}
	tail := map[int]int{} //存后缀的索引
	result := -1
	sum := 0
	for i := n - 1; i >= 0; i-- {
		sum += nums[i]
		tail[sum] = i
		if sum == x && (result > n-i || result == -1) {
			result = n - i
		}
	}
	sum = 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if sum > x {
			break
		} else if sum < x {
			if idx, ok := tail[x-sum]; ok && idx > i &&
				(result == -1 || n-idx+i+1 < result) {
				result = n - idx + i + 1
			}
		} else {
			if result == -1 || i+1 < result {
				result = i + 1
			}
		}
	}
	return result
}

// @lc code=end

/*
优化后缀和为hash表，这样可以快速判断后缀中是否存在前缀中缺少的值

前缀和与后缀和方便计算一段距离的长度和，
然后双指针向中间逼近看长度
Time Limit Exceeded
83/94 cases passed (N/A)
func minOperations(nums []int, x int) int {
	n := len(nums)
	if nums[0] == x || nums[n-1] == x {
		return 1
	}
	prev, tail := make([]int, n), make([]int, n)
	prev[0], tail[n-1] = nums[0], nums[n-1]
	result := -1
	for i := 1; i < n; i++ {
		prev[i] = prev[i-1] + nums[i]
		if prev[i] == x && (result > i+1 || result == -1) {
			result = i + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		tail[i] = nums[i] + tail[i+1]
		if tail[i] == x && (result > n-i || result == -1) {
			result = n - i
		}
	}
	currentLength := 0
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if prev[i]+tail[j] == x {
				currentLength = i + 1 + n - j
				if currentLength < result || result == -1 {
					result = currentLength
				}
				break
			}
		}
	}
	return result
}
*/
