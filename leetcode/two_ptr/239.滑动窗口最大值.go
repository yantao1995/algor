package leetcode

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	result := []int{}
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if i < k {
			if nums[i] > nums[maxIndex] {
				maxIndex = i
			}
			if i == k-1 {
				result = append(result, nums[maxIndex])
			}
		} else {
			if i-k >= maxIndex {
				maxIndex++
				for j := i - k + 1; j <= i; j++ {
					if nums[j] >= nums[maxIndex] {
						maxIndex = j
					}
				}
			} else {
				if nums[i] > nums[maxIndex] {
					maxIndex = i
				}
			}
			result = append(result, nums[maxIndex])
		}
	}
	return result
}

// @lc code=end

/*
	减少k个长度内的遍历次数
	例: 1,4,2,3     k=2
	当i走到1时，长度等于k找到索引为1的数，记录当前索引
	之后i走到2时，只需要判断新加进来的数是不是大于当前记录的索引的数的值，
	分两种情况 :
	a.记录的索引仍处于滑动窗口内:如果小于，那么就取当前记录的索引的值，如果大于，就取新进来的数的值
	b.记录的索引处于滑动窗口外:直接找新窗口内的最大值,并记录最大索引
*/
