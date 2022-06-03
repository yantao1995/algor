package leetcode

/*
 * @lc app=leetcode.cn id=829 lang=golang
 *
 * [829] 连续整数求和
 */

// @lc code=start
func consecutiveNumbersSum(n int) int {
	result := 0
	temp := 0
	for i := 1; i <= n && temp < n; i++ {
		if (n-temp)%i == 0 {
			result++
		}
		temp += i
	}
	return result
}

// @lc code=end

/*
	当n=15时，观察数列
	个数     组合			特征
	1	     15			   15*1+0
	2		 7,8		   7*2+1
	3		 4，5，6		4*3+1+2
	5		 1，2，3，4,5	1*5+1+2+3+4

	所以以第一个数为基数，然后每个数都比前一个大1，所以可以得到推理
	当个数为 m 时，  设a = (n- (1加到m-1) )
	只要 b = a/m ，b为整数，那么就b就是这个组数的第一个数
*/
