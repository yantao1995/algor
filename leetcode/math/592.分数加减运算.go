package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=592 lang=golang
 *
 * [592] 分数加减运算
 */

// @lc code=start
func fractionAddition(expression string) string {
	symbol1 := "+"
	symbol2, temp1, temp2 := "+", "0", "1"
	r1i, r2i, temp1i, temp2i := 0, 1, 0, 0
	i := 0
	if expression[0] != '-' {
		symbol2 = "+"
		temp1 = string(expression[i])
		i++
		if expression[i] == '0' {
			temp1 += "0"
			i++
		}
		i++
		temp2 = string(expression[i])
		i++
		if i < len(expression) && expression[i] == '0' {
			temp2 += "0"
			i++
		}
	}
	calc := func() {
		temp1i, _ = strconv.Atoi(temp1)
		temp2i, _ = strconv.Atoi(temp2)
		if r2i != temp2i {
			r1i, r2i, temp1i, temp2i = r1i*temp2i, r2i*temp2i, temp1i*r2i, temp2i*r2i
		}

		if symbol1 == symbol2 {
			r1i += temp1i
		} else {
			if r1i >= temp1i {
				r1i -= temp1i
			} else {
				r1i = temp1i - r1i
				symbol1 = symbol2
			}
		}

		for temp := 2; temp <= r1i && temp <= r2i; temp++ {
			if r1i%temp == 0 && r2i%temp == 0 {
				r1i /= temp
				r2i /= temp
				temp = 1
			}
		}
		temp1, temp2 = "0", "1"

	}
	calc()
	for i < len(expression) {
		symbol2 = string(expression[i])
		i++
		if i < len(expression) {
			temp1 = string(expression[i])
		}
		i++
		if i < len(expression) && expression[i] == '0' {
			temp1 += "0"
			i++
		}
		i++
		if i < len(expression) {
			temp2 = string(expression[i])
			i++
		}
		if i < len(expression) && expression[i] == '0' {
			temp2 += "0"
			i++
		}
		calc()
	}
	calc()
	if symbol1 == "+" || r1i == 0 {
		symbol1 = ""
		if r1i == 0 {
			r2i = 1
		}
	}
	return symbol1 + strconv.Itoa(r1i) + "/" + strconv.Itoa(r2i)
}

// @lc code=end

/*
	calc 进行运算，然后约分
	for 循环，对分式进行拆解

	初始化默认值为0 ， 即 分子为0，分母为1

	因为每个数字的范围是0-10，所以截取数字的时候，
	需要多判断一位，为了防止越界需要增加判断 i < len(expression)

	calc : r1i,r2i 保存结果集，temp1,temp2,则是截取的分子和分母

	运算时， symbol1 记录结果集的正负， symbol2记录需要计算的式子的正负

*/
