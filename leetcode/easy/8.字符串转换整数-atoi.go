package easy

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	flag, rtn, times, bts := "", 0, 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '-' || str[i] == '+' {
			if bts > 0 {
				break
			}
			times++
			if str[i] == '-' {
				flag = "-"
			}
		} else if str[i] >= '0' && str[i] <= '9' && rtn < 1<<31-1 {
			bts++
			rtn = (rtn * 10) + int(str[i]-'0')
		} else if str[i] == ' ' && rtn == 0 && times == 0 && bts == 0 {
			continue
		} else {
			break
		}
	}

	if times > 1 {
		return 0
	}
	if flag == "-" {
		rtn = 0 - rtn
	}
	if rtn < -1<<31 {
		return -1 << 31
	}
	if rtn > 1<<31-1 {
		return 1<<31 - 1
	}
	return rtn
}

// @lc code=end
