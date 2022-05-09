package leetcode

/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */

// @lc code=start
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1)
	for k := range dp {
		dp[k] = make([]bool, len(s)+1)
		if k == 0 {
			dp[k][0] = true
		} else if p[k-1] == '*' {
			dp[k][0] = dp[k-1][0]
		}
	}
	for i := 0; i < len(p); i++ {
		for j := 0; j < len(s); j++ {
			if dp[i][j] && (p[i] == '?' || p[i] == s[j]) {
				dp[i+1][j+1] = true
			} else if p[i] == '*' {
				dp[i+1][j+1] = dp[i][j+1]
				if dp[i][j] {
					for ; j < len(s); j++ {
						dp[i+1][j+1] = true
					}
					break
				}
			}
		}
	}
	return dp[len(p)][len(s)]
}

// @lc code=end

/*
	动态规划，dp[i][j] 表示 p 从0-i的长度 与 s 从0-j能否匹配
	长度定义 i = len(p)+1   j = len(s)+1 ,因为当前dp[i][j]需要依赖d[i-1][j-1]
		为了计算时忽略边界值考虑，所以给边界拓宽了1，dp[i][j]映射到了dp[i+1][j+1]
	需要初始化dp数组的第一列,因为第一个字符需要判断是否匹配，所以dp[0][0] 初始化为true,又因为'*'可以匹配0个
		所以如果上个字符是 '*' 那么当前字符能否匹配应该和上一个字符一样 即 dp[k][0] = dp[k-1][0]

	初始化完成后的p[i]=='*'时， 也是同理 dp[i+1][j+1] = dp[i][j+1]

	条件1 :  p[i] == '?' || p[i] == s[j] 时，表示当前 两个字符可以匹配，
		即dp[i+1][j+1]为true,但是前提条件就是p[i-1]与s[j-1]可以匹配，即 dp[i][j]为true

	条件2 :  p[i] == '*' ,表示当前字符可以匹配到任意长度，
		如果匹配0个，那么在匹配到 s[j] 个长度时候 p[i] 与p[i-1]的效果应该是一样的，即: dp[i+1][j+1] = dp[i][j+1]
		如果匹配大于0个，那么只要 p[i-1]与 s[j-1]能够匹配，那么后面所有的字符都可以匹配，
			即dp[i][j]，那么dp[i+1][j+1]到dp[i+1][末尾] 都可以匹配,然后跳出当前for循环，减少遍历次数

	例子 ：  s="abceb"  p="*a*b"

	   0	 a		b		c		e		b
	0 true   false  false   false   false   false
	* true   true   true    true    true    true
	a false  true   false   false   false   false
	* false  true   true    true    true    true
	b false  false  true    false   false   true

*/
