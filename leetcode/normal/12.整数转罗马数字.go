package leetcode

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	rtn := ""
	symbolMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	seqKey := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	for i := len(seqKey) - 1; i >= 0; i-- {
		count := 0
		if num < seqKey[i] {
			continue
		}
		count = num / seqKey[i]
		num = num - count*seqKey[i]
		for count > 0 {
			rtn += symbolMap[seqKey[i]]
			count--
		}

	}
	return rtn
}

// @lc code=end
