package leetcode

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	total := "0"
	for i := len(num1); i >= 0; i-- {
		for j := len(num2); j >= 0; j-- {

		}
	}
	return total
}

func strAdd(a, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	// := []byte{}
	for i := len(a); i >= 0; i-- {
		weight := '0' - 48
		for j := len(b); j >= 0; j-- {
			temp := a[i] + b[j] - 96
			temp += weight
			if temp >= 10 {
				weight = '1' - 48
			} else {
				weight = 0
			}

		}

	}
	return ""
}

// @lc code=end
