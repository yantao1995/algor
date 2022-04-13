package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=93 lang=golang
 *
 * [93] 复原 IP 地址
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	result := []string{}
	isLegal := func(i, j int) bool { // [i,j] 闭区间
		if j > i {
			if s[i] == '0' {
				return false
			}
			sInt, _ := strconv.Atoi(s[i : j+1])
			if sInt > 255 {
				return false
			}
		}
		return true
	}
	bs := []byte(s)
	var backtrace func(i, count int, bts []byte)
	backtrace = func(i, count int, bts []byte) {
		if i == len(bs) || count > 3 {
			if len(bts) == cap(bts) {
				result = append(result, string(append([]byte{}, bts...)))
			}
			return
		}
		if len(bts) > 0 {
			bts = append(bts, '.')
		}
		for ; i < len(bs); i++ {
			for j := i; j < len(bs) && j-i < 3; j++ {
				if isLegal(i, j) {
					backtrace(j+1, count+1, append(bts, append([]byte{}, bs[i:j+1]...)...))
				}
			}
		}

	}
	backtrace(0, 0, make([]byte, 0, len(bs)+3))
	return result
}

// @lc code=end

/*
	isLegal 用户判断该字串是否合法
	然后用count记录 . 添加了多少个，如果 .添加的个数 大于3，
	那么ip地址就生成完成了，然后再检测长度是否为预期，符合预期就是正确的ip地址

	剪枝 :
		1. count >3 ，大于3个 . 的地址永远不合法
		2.  [i,j] 区间长度大于3  ，255是最大的数，长度为 3

*/
