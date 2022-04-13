package leetcode

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	result := [][]string{}
	dp := make([][]bool, len(s)) //标识回文串  dp[i][j] 标识 [i,j]是回文串
	for k := range dp {
		dp[k] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if i == j || ((j-i == 1 || dp[i+1][j-1]) && s[i] == s[j]) {
				dp[i][j] = true
			}
		}
	}
	strs := make([]string, 0, len(s))
	var backtrace func(i, count int, strs []string)
	backtrace = func(i, count int, strs []string) {
		if i == len(s) || i != count {
			if count == len(s) {
				result = append(result, append([]string{}, strs...))
			}
			return
		}
		for ; i < len(s); i++ {
			for j := i; j < len(s); j++ {
				if dp[i][j] {
					backtrace(j+1, count+(j+1-i), append(strs, s[i:j+1]))
				}
			}
		}
	}
	backtrace(0, 0, strs)
	return result
}

// @lc code=end

/*
	先利用动归生成dp数组，dp[i][j]表示区间 [i,j] 为回文串
	然后利于回溯挨个找回文串,然后添加进当前记录string数组strs
	使用i来标识当前走过的路径，然后使用count来标识当前 strs 内保存的字符串个数，
	i != count 用于剪去当前不符合要求的结果集
*/
