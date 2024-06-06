package lcpr

/*
 * @lc app=leetcode.cn id=1673 lang=golang
 * @lcpr version=30202
 *
 * [1673] 找出最具竞争力的子序列
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func mostCompetitive(nums []int, k int) []int {
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		//len(nums)-i > k-len(stack) 加上当前的 i，应该刚好能放进去
		for len(stack) > 0 && nums[i] < stack[len(stack)-1] && len(nums)-i > k-len(stack) {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < k {
			stack = append(stack, nums[i])
		}
	}
	return stack
}

// @lc code=end

/*
// @lcpr case=start
// [3,5,2,6]\n2\n
// @lcpr case=end

// @lcpr case=start
// [2,4,3,3,5,4,9,6]\n4\n
// @lcpr case=end

*/

/*
优化：每次贪当前最小的值即可。


超时 85/88
每次拿出剩余数列中的最小值，前提是若选中当前的最小值，并且加上后面的全部数的数量必须满足k的个数
【bug:如果数据全部一样，那么每次都扫到同一个数】
func mostCompetitive(nums []int, k int) []int {
	result := make([]int, k)
	minIndex := 0
	for i := 0; i < len(result); i++ {
		for j := minIndex; j+k-i <= len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		result[i] = nums[minIndex]
		minIndex++
	}
	return result
}
*/
