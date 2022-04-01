package leetcode

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) []int {
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			nums[i] = nums[i/2]
		} else {
			nums[i] = nums[i-1] + 1
		}
	}
	return nums
}

// @lc code=end

/*
	对于2的倍数，即
	0 --> 0
	1 --> 1
	2 --> 10
	3 --> 11
	4 --> 100
	5 --> 101
	6 --> 110

	分偶数与奇数
	偶数。比如：6。6是3的两倍，即3右移一位，3<<1。那么1的数量相等
	奇数。比如：3。3是2+1 ，即在2的基础上，加了1位。
*/
