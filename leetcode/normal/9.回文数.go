package leetcode

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
// func isPalindrome(x int) bool { //不转字符串
// 	if x < 0 {
// 		return false
// 	}
// 	if x < 10 {
// 		return true
// 	}
// 	oldX, crease, new := x, 10, 0
// 	isTen := false
// 	temp := x
// 	for temp > 1 { //看是不是10000这种  还是 98000这种，首位为1，其余位为0
// 		temp = temp / 10
// 	}
// 	if temp == 1 { //是10000这种
// 		isTen = true
// 	}
// 	for {
// 		if crease/x > 10 && isTen {
// 			x = new
// 			break
// 		}
// 		if crease/x >= 10 && !isTen {
// 			x = new
// 			break
// 		}
// 		if new > 0 {
// 			new *= 10
// 		}
// 		new += (x % crease / (crease / 10))
// 		crease *= 10
// 	}
// 	return x == oldX
// }

// @lc code=end
