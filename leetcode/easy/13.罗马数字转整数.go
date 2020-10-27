package easy

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	rtn := 0
	symbolMap := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	seqKey := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	for i := 0; i < len(s); {
		for j := 2; j > 0; j-- {
			for m := len(seqKey) - 1; m >= 0; m-- {
				if i+j > len(s) {
					break
				}
				if s[i:i+j] == seqKey[m] {
					rtn += symbolMap[seqKey[m]]
					i += j
					goto comeHere
				}
			}
		}
	comeHere:
	}
	return rtn
}

// @lc code=end
