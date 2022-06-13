package leetcode

/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] 猜数字大小
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;

 */ //

// func guessNumber(n int) int {
// 	last := 0
// 	result := guess(n)
// 	for result != 0 {
// 		if result == -1 {
// 			last = n
// 			n /= 2
// 		} else {
// 			n = (n + last) / 2
// 		}
// 		result = guess(n)
// 	}
// 	return n
// }

// @lc code=end
