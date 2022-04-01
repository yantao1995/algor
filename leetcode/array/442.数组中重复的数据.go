package leetcode

/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	result := []int{}
	n := len(nums)
	for k := range nums {
		temp := nums[k] - 1
		for temp >= n {
			temp -= n
		}
		nums[temp] += n
	}
	for k := range nums {
		if nums[k] > 2*n && nums[k] <= 3*n {
			result = append(result, k+1)
		}
	}
	return result
}

// @lc code=end

/*
	遍历元素，让元素的值变成下次遍历的索引下标，并给索引下标的所在值+n
	因为范围是 [1,n]，如果存在两个 3， 那么索引为3的元素必然会增加两次
	遍历的时候，重复的值都是对应的2倍到3倍之间
*/
