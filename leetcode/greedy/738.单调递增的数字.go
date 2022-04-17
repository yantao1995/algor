package leetcode

/*
 * @lc app=leetcode.cn id=738 lang=golang
 *
 * [738] 单调递增的数字
 */

// @lc code=start
func monotoneIncreasingDigits(n int) int {
	source := []int{}
	for n > 0 {
		source = append(source, n%10)
		n /= 10
	}
	for i := 0; i < len(source); i++ {
		if i+1 < len(source) && source[i] < source[i+1] {
			for j := i; j >= 0; j-- {
				source[j] = 9
			}
			source[i+1]--
		}
	}
	result := 0
	for i, weight := 0, 1; i < len(source); i, weight = i+1, weight*10 {
		result += source[i] * weight
	}
	return result
}

// @lc code=end

/*
	每次都贪当前位取最大值，因为可以相等，即当前位如果比后一位小，直接取9
	比如 1000
	source里面的值是 0001
	走到索引为2时，判断 0比 1小
	那么直接索引j向前全部变成9，
	索引3的值自减1
*/
