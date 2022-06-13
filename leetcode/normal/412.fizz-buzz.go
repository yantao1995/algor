package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	temp := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			temp[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			temp[i-1] = "Fizz"
		} else if i%5 == 0 {
			temp[i-1] = "Buzz"
		} else {
			temp[i-1] = strconv.Itoa(i)
		}
	}
	return temp
}

// @lc code=end
