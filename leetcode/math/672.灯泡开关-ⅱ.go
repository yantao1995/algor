package leetcode

/*
 * @lc app=leetcode.cn id=672 lang=golang
 *
 * [672] 灯泡开关 Ⅱ
 */

// @lc code=start
func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	if n == 1 {
		return 2
	} else if n == 2 {
		if presses == 1 {
			return 3
		}
		return 4
	} else {
		if presses > 2 {
			return 8
		} else {
			if presses == 1 {
				return 4
			}
			return 7
		}
	}
}

// @lc code=end

/*
	n大于3的时候，press>2时，结果都是8，其他的可以依次枚举结果
*/
