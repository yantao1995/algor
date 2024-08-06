package lcpr

import (
	"strconv"
)

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
	calc := func(temp []int) int {
		for i := 1; i < len(temp); i += 2 {
			if temp[i] == -1 {
				temp[i-1] *= temp[i+1]
				temp = append(temp[:i], temp[i+2:]...)
				i -= 2
			}
		}
		for len(temp) > 1 {
			if temp[1] == -2 {
				temp[2] += temp[0]
			} else {
				temp[2] = temp[0] - temp[2]
			}
			temp = temp[2:]
		}
		return temp[0]
	}

	sign := []int{-1, -2, -3}
	// -1 x  -2 +  -3 -
	var backtrace func(i int, temp []int)
	backtrace = func(i int, temp []int) {
		if i == len(num) {
			bts := make([]byte, 0, len(temp))
			for _, v := range temp {
				if v == -1 {
					bts = append(bts, '*')
				} else if v == -2 {
					bts = append(bts, '+')
				} else if v == -3 {
					bts = append(bts, '-')
				} else {
					bts = append(bts, []byte(strconv.Itoa(v))...)
				}
			}

			if calc(append([]int(nil), temp...)) == target {
				result = append(result, string(bts))
			}
			return
		}
		for j := 0; j < len(sign); j++ {
			for m, last := i, 0; m < len(nums) && (m == i || nums[i] > 0); m++ {
				last = last*10 + nums[m]
				temp = append(temp, sign[j], last)
				backtrace(m+1, temp)
				temp = temp[:len(temp)-2]
			}
		}
	}

	for i, last := 0, 0; i < len(nums) && (i == 0 || nums[0] > 0); i++ {
		last = last*10 + nums[i]
		backtrace(i+1, []int{last})
	}
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

/*
	枚举所有的情况，然后计算结果即可。
	需注意首数字不为0的话，可以将1-n个数字合并成一个大数字
*/
