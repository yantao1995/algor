package easy

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	if x < 10 && x > -10 {
		return x
	}
	crease, new := 10, 0
	flag, isTen := false, false
	if x < 0 {
		x = 0 - x
		flag = true
	}
	temp := x
	for temp > 1 { //看是不是10000这种  还是 98000这种，首位为1，其余位为0
		temp = temp / 10
	}
	if temp == 1 { //是10000这种
		isTen = true
	}
	for {
		if crease/x > 10 && isTen {
			x = new
			break
		}
		if crease/x >= 10 && !isTen {
			x = new
			break
		}
		if new > 0 {
			new *= 10
		}
		new += (x % crease / (crease / 10))
		crease *= 10
	}
	if flag {
		x = 0 - x
	}
	if x > 1<<31-1 || x < -1<<31 {
		return 0
	}
	return x
}

// @lc code=end
