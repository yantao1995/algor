package leetcode

/*
 * @lc app=leetcode.cn id=2180 lang=golang
 *
 * [2180] 统计各位数字之和为偶数的整数个数
 */

// @lc code=start
func countEven(num int) int {
	result := 0
	for i := 1; i <= num; i++ {
		sum := 0
		for temp := i; temp > 0; temp /= 10 {
			sum += temp % 10
		}
		if sum&1 == 0 {
			result++
		}
	}
	return result
}

// @lc code=end

/*
	直接挨个对每个数字的位数累加判断即可
*/
