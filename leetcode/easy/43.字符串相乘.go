package easy

import "fmt"

/*
 * @lc app=leetcode.cn id=43 lang=golang
 *
 * [43] 字符串相乘
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	for k := 0; k < len(num1); k++ { //去首位0 num1
		if num1[k] == '0' {
			num1 = num1[k+1:]
			k--
		} else {
			break
		}
	}
	for k := 0; k < len(num2); k++ { //去首位0 num2
		if num2[k] == '0' {
			num2 = num2[k+1:]
			k--
		} else {
			break
		}
	}
	total := []byte{'0'}
	zero := []byte{'0'}
	result := []byte{}
	weight := zero[0] - 48
	totalWeight := []byte{}
	n1, n2 := []byte(num1), []byte(num2)
	var a, b, temp byte
	for i := len(n1) - 1; i >= 0; i-- {
		weight = zero[0] - 48
		result = []byte{}
		a = n1[i] - 48
		for j := len(n2) - 1; j >= 0; j-- {
			b = n2[j] - 48
			temp = a * b
			temp += weight
			if temp >= 10 {
				weight = temp / 10
			} else {
				weight = 0
			}
			bit := []byte{temp%10 + 48}
			if j == 0 {
				bit = append([]byte{weight + 48}, bit...)
			}
			result = append(bit, result...)
		}
		if len(totalWeight) > 0 {
			result = append(result, totalWeight...)
		}
		fmt.Println("total, result", string(total), string(result))
		total = strAdd(total, result)
		totalWeight = append(totalWeight, '0')
	}
	if len(total) == 0 {
		return "0"
	}
	return string(total)
}

func strAdd(a, b []byte) []byte {
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
