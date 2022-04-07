package leetcode

/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
func longestValidParentheses(s string) int {
	max := 0
	dp := make([]int, len(s))
	stack := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
				dp[i] = dp[i-1] + 2
				if i-dp[i] > 0 {
					dp[i] += dp[i-dp[i]]
				}
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
	}

	return max
}

// @lc code=end

/*
	max 记录当前所找到的最大的长度
	dp[i] 记录到当前位置,能凑成的最大长度
	stack 用户匹配左右括号 (),匹配成功时，就设置dp数组的值

	dp[i] = dp[i-1] + 2  匹配成功一个括号对时，当前的长度应该 +2

	dp[i] += dp[i-dp[i]] 当匹配成功时，需要去找当前括号对前面的一个 dp的最大括号长度,
	  为了实现  ()() 像这种括号，在遍历到最后一个位置的时候即 dp[3]=2 ，但是需要加上前面的
	  有效括号长度即使， dp[3] += dp[3-dp[2]] ,减去当前dp[i]长度就是括号匹配的前一个位置
	  因为左括号 (  的dp 长度为0，所以不需要额外判断

	最后将每次得到的 dp[i] 与max 得到当前的最大值
*/
