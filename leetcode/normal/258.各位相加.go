package leetcode

/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	weight, temp := 1, 0
	for num >= 10 {
		temp, weight = 0, 1
		for weight <= num {
			weight *= 10
		}
		weight /= 10
		for weight >= 1 {
			temp += (num / weight)
			num %= weight
			weight /= 10
		}
		num = temp
	}
	return num
}
func strAdd258(a, b []byte) []byte {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
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
			return result
		}
	}
	if len(result) == 0 {
		return []byte{'0'}
	}
	return result
}

// @lc code=end
