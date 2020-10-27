package easy

/*
 * @lc app=leetcode.cn id=371 lang=golang
 *
 * [371] 两整数之和
 */

// @lc code=start

//无进位加法  a^b
//进位加法   (a & b) << 1

func getSum(a int, b int) int {
	if b == 0 {
		return a
	}
	sum := a ^ b
	bit := (a & b) << 1
	return getSum(sum, bit)
}

// @lc code=end
