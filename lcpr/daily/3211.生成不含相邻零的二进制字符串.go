package lcpr

/*
 * @lc app=leetcode.cn id=3211 lang=golang
 * @lcpr version=30204
 *
 * [3211] 生成不含相邻零的二进制字符串
 */

// @lcpr-template-start

// @lcpr-template-end
// @lc code=start
func validStrings(n int) []string {
	result := []string{}
	var backtrace func(str string)
	backtrace = func(str string) {
		if len(str) == n {
			result = append(result, str)
			return
		}
		backtrace(str + "1")
		if str[len(str)-1] == '1' {
			backtrace(str + "0")
		}
	}
	backtrace("0")
	backtrace("1")
	return result
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/

/*

优化思路，把枚举变成递归，代码逻辑清晰



//直接枚举每种情况
func validStrings(n int) []string {
	result := []string{"0", "1"}
	for i := 1; i < n; i++ {
		temp := make([]string, 0, len(result)*2)
		for _, v := range result {
			temp = append(temp, v+"1")
			if v[len(v)-1] == '1' {
				temp = append(temp, v+"0")
			}
		}
		result = temp
	}
	return result
}
*/
