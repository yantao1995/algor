package leetcode

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=753 lang=golang
 *
 * [753] 破解保险箱
 */

// @lc code=start
func crackSafe(n int, k int) string {
	seen := map[int]bool{}
	ans := ""
	highest := int(math.Pow(10, float64(n-1)))

	var dfs func(int)
	dfs = func(node int) {
		for x := 0; x < k; x++ {
			nei := node*10 + x
			if !seen[nei] {
				seen[nei] = true
				dfs(nei % highest)
				ans += strconv.Itoa(x)
			}
		}
	}
	dfs(0)
	for i := 1; i < n; i++ {
		ans += "0"
	}
	return ans
}

// @lc code=end

/*

参考官方题解，欧拉回路 Hierholzer 算法


两次回溯枚举所有组合 超时
第一次回溯枚举答案
第二次回溯枚举组合数

Time Limit Exceeded
	13/38 cases passed (N/A)
Testcase
	2
	4
Expected Answer
	"03322312113020100"

func crackSafe(n int, k int) string {
	ans := []string{}
	var backtrace func(str string)
	backtrace = func(str string) {
		if len(str) == n {
			ans = append(ans, str)
			return
		}
		for i := '0'; i < rune('0'+k); i++ {
			backtrace(str + string(i))
		}
	}
	backtrace("")
	combine := func(str1, str2 string) string {
		i := len(str2)
		for i > 0 {
			if strings.HasSuffix(str1, str2[:i]) {
				break
			}
			i--
		}
		return str1 + str2[i:]
	}
	result := ""
	distinct := map[int]bool{}
	var backtrace2 func(count int, temp string)
	backtrace2 = func(count int, temp string) {
		if result != "" && len(temp) >= len(result) {
			return
		}
		if count == len(ans) {
			if result == "" || len(result) > len(temp) {
				result = temp
			}
			return
		}
		for i := 0; i < len(ans); i++ {
			if !distinct[i] {
				distinct[i] = true
				backtrace2(count+1, combine(temp, ans[i]))
				distinct[i] = false
			}
		}
	}
	backtrace2(0, "")
	return result
}
*/
