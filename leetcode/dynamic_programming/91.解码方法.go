package leetcode

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	s = " " + s
	dp := make([]int, len(s))
	dp[0], dp[1] = 1, 1
	for i := 2; i < len(s); i++ {
		dp[i] = dp[i-1]
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i] = dp[i-2]
				continue
			}
			return 0
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7' && s[i] > '0') {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(dp)-1]
}

// @lc code=end

/*
	dp[i] 表示到 0-i 时，有多少种数量
	为了方便计算 dp[i-2]，所以直接给s前面增加了一个空格,然后从i=2开始计算
		就相当于从 s[1]开始计算，然后开始直接判断 s[0]，如果为0就直接退出
		所以就保证了 dp[0]与dp[1]可以直接初始化为1

	条件1  s[i] =='0',如果前面一个也是 '0'，那么直接return ，否则dp[i]应该与dp[i-2]一致
			因为相当于 s[i-1]被 s[i]的'0'牵制了，无法与前面的数进行组合。
	条件2  s[i-1]=='1'  || (s[i-1] == '2' && s[i] < '7' && s[i] > '0') 表示当前不为0
			并且可以与前一个字符组成一个组合，那么与前一个字符组合即dp[i-2]个数，
			不与前一个字符组合即dp[i-1] + dp[i-2]个数
	条件3  默认条件，就是当前字符不为 '0' 且前一个字符不为 '1' 和 '2'，
			那么当前字符只能为单字符，所以个数应该等于dp[i-1]
*/
