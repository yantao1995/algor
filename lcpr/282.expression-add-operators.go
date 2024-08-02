package lcpr

/*
 * @lc app=leetcode.cn id=282 lang=golang
 * @lcpr version=30204
 *
 * [282] 给表达式添加运算符
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func addOperators(num string, target int) []string {
	result := []string{}
	nums := make([]int, len(num))
	for i := 0; i < len(num); i++ {
		nums[i] = int(num[i] - '0')
	}
	var dac func(temp []int, i, j int) int
	dac = func(temp []int, i, j int) int {
		if i == j {
			return temp[i]
		}
		for k := 1; k < len(temp); k += 2 {
			if temp[k] == -1 {
				temp[k-1] *= temp[k+1]
				temp = append(temp[:k], temp[k+2:]...)
				return dac(0)
			}
		}
		if temp[1] == -2 {
			temp[2] += temp[0]
		} else {
			temp[2] -= temp[0]
		}
		return dac(temp, 2, j)
	}

	check := func(temp []int) {

	}
	sign := []int{-1, -2, -3}
	// -1 x  -2 +  -3 -
	var backtrace func(i int, temp []int)
	backtrace = func(i int, temp []int) {
		if i == len(num) {
			check(temp)
			return
		}
		for j := 0; j < len(sign); j++ {
			temp = append(temp, sign[j], nums[i])
			backtrace(i+1, temp)
			temp = temp[:len(temp)-2]
		}
	}
	backtrace(1, []int{nums[0]})
	return result
}

// @lc code=end

/*
// @lcpr case=start
// "123"\n6\n
// @lcpr case=end

// @lcpr case=start
// "232"\n8\n
// @lcpr case=end

// @lcpr case=start
// "3456237490"\n9191\n
// @lcpr case=end

*/
