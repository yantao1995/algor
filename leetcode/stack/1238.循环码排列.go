package leetcode

/*
 * @lc app=leetcode.cn id=1238 lang=golang
 *
 * [1238] 循环码排列
 */

// @lc code=start
func circularPermutation(n int, start int) []int {
	stack := make([]int, 0, 2<<n-1)
	stack = append(stack, 0)
	for i := 1; i <= n; i++ { //在第几位加权1  ，右边第一个为1
		for tail := len(stack) - 1; tail >= 0; tail-- {
			stack = append(stack, stack[tail]+1<<(i-1))
		}
	}
	for k := range stack {
		if stack[k] == start {
			return append(stack[k:], stack[:k]...)
		}
	}
	return stack
}

// @lc code=end

/*
	格雷编码，最后找到 start位置，当第一个即可
*/
