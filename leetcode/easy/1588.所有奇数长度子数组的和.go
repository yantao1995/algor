package easy

/*
 * @lc app=leetcode.cn id=1588 lang=golang
 *
 * [1588] 所有奇数长度子数组的和
 */

// @lc code=start
func sumOddLengthSubarrays(arr []int) int {
	ht := map[int]int{}
	result := 0
	length := 0
	for k := range arr {
		length = 0
		for k+length < len(arr) {
			for m := k; m <= k+length; m++ {
				ht[m]++
			}
			length += 2
		}
	}
	for k := range ht {
		result += ht[k] * arr[k]
	}
	return result
}

// @lc code=end
