package leetcode

/*
 * @lc app=leetcode.cn id=2488 lang=golang
 *
 * [2488] 统计中位数为 K 的子数组
 */

// @lc code=start
func countSubarrays(nums []int, k int) int {
	l, r := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			l, r = i-1, i+1
			break
		}
	}
	result := 1
	l0, r0 := 0, 0
	leftCount, greatCount := make([]int, l+1), make([]int, len(nums)-r)
	i := 0
	lm, rm := map[int]int{}, map[int]int{}
	for l >= 0 {
		leftCount[i] = -1
		if nums[l] > k {
			leftCount[i] = 1
		}
		if i > 0 {
			leftCount[i] += leftCount[i-1]
		}
		if leftCount[i] == 0 || leftCount[i] == 1 {
			result++
		}
		lm[leftCount[i]]++
		i++
		l--
	}
	i = 0
	for r < len(nums) {
		greatCount[i] = -1
		if nums[r] > k {
			greatCount[i] = 1
		}
		if i > 0 {
			greatCount[i] += greatCount[i-1]
		}
		if greatCount[i] == 0 || greatCount[i] == 1 {
			result++
		}
		rm[greatCount[i]]++
		i++
		r++
	}
	result += l0 + r0
	for ll, lv := range lm {
		if rm[-ll] > 0 {
			result += lv * rm[-ll]
		}
		if rm[1-ll] > 0 {
			result += lv * rm[1-ll]
		}
	}
	return result
}

// @lc code=end

/*

优化,降低左右计算的次数，用map存左右，如果当前rm中存在与lm中
相加 ==0 或 == -1 的值，那么左右相乘即为组合的个数


超时，记录左右两边从k位置向两头的计数(大于k和小于k相减的个数)
Time Limit Exceeded
43/45 cases passed (N/A)
func countSubarrays(nums []int, k int) int {
	l, r := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			l, r = i-1, i+1
			break
		}
	}
	result := 1
	l0, r0 := 0, 0
	leftCount, greatCount := make([]int, l+1), make([]int, len(nums)-r)
	i := 0
	for l >= 0 {
		leftCount[i] = -1
		if nums[l] > k {
			leftCount[i] = 1
		}
		if i > 0 {
			leftCount[i] += leftCount[i-1]
		}
		if leftCount[i] == 0 || leftCount[i] == 1 {
			result++
		}
		i++
		l--
	}
	i = 0
	for r < len(nums) {
		greatCount[i] = -1
		if nums[r] > k {
			greatCount[i] = 1
		}
		if i > 0 {
			greatCount[i] += greatCount[i-1]
		}
		if greatCount[i] == 0 || greatCount[i] == 1 {
			result++
		}
		i++
		r++
	}
	result += l0 + r0
	for i := 0; i < len(leftCount); i++ {
		for j := 0; j < len(greatCount); j++ {
			if leftCount[i]+greatCount[j] == 0 || leftCount[i]+greatCount[j] == 1 {
				result++
			}
		}
	}
	return result
}
*/
