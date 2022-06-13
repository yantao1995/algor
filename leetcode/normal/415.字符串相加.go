package leetcode

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {
	a, b := []byte(num1), []byte(num2)
	if len(a) == 0 {
		return num2
	}
	if len(b) == 0 {
		return num1
	}
	zero := []byte{'0'}
	result := []byte{}
	weight := zero[0] - 48
	for index := 1; index <= len(a) || index <= len(b); index++ {
		var temp byte
		if index <= len(a) {
			temp += (a[len(a)-index] - 48)
		}
		if index <= len(b) {
			temp += (b[len(b)-index] - 48)
		}
		temp += weight
		if temp >= 10 {
			weight = 1
		} else {
			weight = 0
		}
		bit := []byte{temp%10 + 48}
		result = append(bit, result...)
	}
	if weight > 0 {
		result = append([]byte{'1'}, result...)
	}
	for k := 0; k < len(result); k++ { //去首位0
		if result[k] == '0' {
			result = result[k+1:]
			k--
		} else {
			return string(result)
		}
	}
	if len(result) == 0 {
		return "0"
	}
	return string(result)
}

// @lc code=end
