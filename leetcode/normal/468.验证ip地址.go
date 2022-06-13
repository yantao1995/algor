package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */

// @lc code=start
func validIPAddress(queryIP string) string {
	slice := strings.Split(queryIP, ".")
	if len(slice) == 4 {
		for k := range slice {
			i, err := strconv.Atoi(slice[k])
			if err != nil || strconv.Itoa(i) != slice[k] ||
				i < 0 || i > 255 {
				return "Neither"
			}
		}
		return "IPv4"
	}
	slice = strings.Split(queryIP, ":")
	if len(slice) == 8 {
		for k := range slice {
			if len(slice[k]) < 1 || len(slice[k]) > 4 {
				return "Neither"
			}
			for kk := range slice[k] {
				if (slice[k][kk] >= '0' && slice[k][kk] <= '9') ||
					(slice[k][kk] >= 'a' && slice[k][kk] <= 'f') ||
					(slice[k][kk] >= 'A' && slice[k][kk] <= 'F') {
					continue
				}
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}

// @lc code=end

/*
	先按ipv4分隔字符串，如果长度等于4，那么就进入ipv4的判断，因为ipv6中不包含ipv4的分隔符.
	所以分隔之后，要么是ipv4,要么是ipv6
	ipv4的判断：因为不包含前导0，所以转整型之后再转回去还相等就不包含前导0，然后在做一个大小的判断即可
	ipv6的判断：直接就挨个判断每个字符是不是16进制
*/
