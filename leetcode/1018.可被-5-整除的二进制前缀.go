package leetcode

/*
 * @lc app=leetcode.cn id=1018 lang=golang
 *
 * [1018] 可被 5 整除的二进制前缀
 */

// @lc code=start
func prefixesDivBy5(A []int) []bool {
	reslt := make([]bool, len(A))
	total := 0
	for i := 0; i < len(A); i++ {
		total *= 2
		total %= 5 //降维  防止溢出变成负数
		total += A[i]
		reslt[i] = total%5 == 0
	}
	return reslt
}

// @lc code=end
