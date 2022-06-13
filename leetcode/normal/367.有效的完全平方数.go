package leetcode

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num < 0 {
		return false
	}
	temp := num
	for {
		if temp*temp == num {
			return true
		}
		if temp*temp > num {
			temp /= 2
			continue
		} else {
			if (temp+1)*(temp+1) <= num {
				temp++
			} else {
				return false
			}
		}
	}

}

// @lc code=end
